package jwt

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/pkg/crypto"

	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/metadata"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	json "github.com/json-iterator/go"
)

type JWT interface {
	//Parse(ctx context.Context, dst any) error
	//
	//ParseMapClaims(ctx context.Context) (map[string]any, error)
	//
	//GetToken(ctx context.Context) string
	//SetToken(ctx context.Context, token string) context.Context

	// SetInfo(req *http.Request, data string)
	// GetInfo(ctx context.Context, dst any) error

	ParseByToken(token string, dst any) error

	NewToken(m map[string]any) (string, error)

	IsBlacklist(token string) bool

	AddBlacklist(token string)

	Renewal(ctx context.Context) (string, error)

	SetUnique(key, token string) error

	CompareUniqueToken(key, token string) bool

	Config() *Config
}

type jwt struct {
	conf  *Config
	redis *redis.Client
	rw    sync.RWMutex
}

type tokenKey struct{}

const (
	blackPrefix  = "token_black"
	uniquePrefix = "token_unique"
	infoMdKey    = "x-md-global-info"
)

type Config struct {
	Whitelist      map[string]bool
	Expire         time.Duration
	Renewal        time.Duration
	Secret         string
	App            string
	UniqueDevice   bool
	UniquePlatform bool
}

// New 初始化
func New(ctx kratosx.Context, conf *Config) JWT {
	return &jwt{conf: conf, redis: ctx.Redis()}
}

func (j *jwt) Config() *Config {
	return j.conf
}

// NewToken is create jwt []byte
func (j *jwt) NewToken(m map[string]any) (string, error) {
	if j == nil {
		return "", errors.New("jwt config not enable or configure")
	}

	m["exp"] = jwtv5.NewNumericDate(time.Now().Add(j.conf.Expire + time.Second)) // 过期时间
	m["nbf"] = jwtv5.NewNumericDate(time.Now())                                  // 生效时间
	m["iat"] = jwtv5.NewNumericDate(time.Now())                                  // 签发时间

	keyFunc := func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	}

	jwtToken := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, jwtv5.MapClaims(m))
	key, err := keyFunc(jwtToken)
	if err != nil {
		return "", err
	}

	token, err := jwtToken.SignedString(key)
	if err != nil {
		return "", err
	}

	// 在尾部追加app信息
	token = fmt.Sprintf("%s.%s", token, base64.StdEncoding.EncodeToString([]byte(j.conf.App)))

	return token, nil
}

// Parse 解析token
func (j *jwt) Parse(ctx context.Context, dst any) error {
	claims, err := j.ParseMapClaims(ctx)
	if err != nil {
		return err
	}
	body, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

// ParseByToken 解析指定的token
func (j *jwt) ParseByToken(token string, dst any) error {
	token = token[:strings.LastIndex(token, ".")]
	tokenInfo, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	})
	if tokenInfo == nil || tokenInfo.Claims == nil {
		return errors.New("token parse error")
	}

	claims, is := tokenInfo.Claims.(jwtv5.MapClaims)
	if !is {
		return errors.New("token format error")
	}

	body, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

// ParseMapClaims 解析token到map
func (j *jwt) ParseMapClaims(ctx context.Context) (map[string]any, error) {
	tokenInfo, is := kratosJwt.FromContext(ctx)
	if !is {
		if j == nil {
			return nil, kratosJwt.ErrMissingJwtToken
		}
		token := j.GetToken(ctx)
		if token == "" {
			return nil, kratosJwt.ErrMissingJwtToken
		}

		parser, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
			return []byte(j.conf.Secret), nil
		})
		if parser == nil || parser.Claims == nil {
			return nil, kratosJwt.ErrTokenInvalid
		}

		tokenInfo, is = parser.Claims.(jwtv5.MapClaims)
		if !is {
			return nil, kratosJwt.ErrTokenParseFail
		}
	}
	claims, is := tokenInfo.(jwtv5.MapClaims)
	if !is {
		return nil, kratosJwt.ErrTokenParseFail
	}
	return claims, nil
}

// GetToken 从ctx中获取token
func (j *jwt) GetToken(ctx context.Context) string {
	token, _ := ctx.Value(tokenKey{}).(string)
	return token
}

// SetToken 设置token的值到当前的ctx
func (j *jwt) SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

// Renewal token续期
func (j *jwt) Renewal(ctx context.Context) (string, error) {
	token := j.GetToken(ctx)
	if token == "" {
		return "", errors.New("token is miss")
	}

	tokenInfo, _ := jwtv5.Parse(token, func(token *jwtv5.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	})
	if tokenInfo == nil || tokenInfo.Claims == nil {
		return "", errors.New("token parse error")
	}

	claims, is := tokenInfo.Claims.(jwtv5.MapClaims)
	if !is {
		return "", errors.New("token format error")
	}

	// 判断token失效是否超过10s
	exp := int64(claims["exp"].(float64))
	now := time.Now().Unix()
	if exp > now {
		return "", errors.New("token is alive")
	}

	if now-exp > int64(j.conf.Renewal.Seconds()) {
		return "", errors.New("token is over max renewal time")
	}

	return j.NewToken(claims)
}

func (a *jwt) path(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

// IsWhitelist 判断请求的接口是否在白名单内
func (j *jwt) IsWhitelist(path, method string) bool {
	j.rw.RLock()
	defer j.rw.RUnlock()

	path = j.path(method, path)
	if j.conf.Whitelist[path] {
		return true
	}

	for p := range j.conf.Whitelist {
		// 将*替换为匹配任意多字符的正则表达式
		pattern := "^" + p + "$"
		pattern = regexp.MustCompile(`\*`).ReplaceAllString(pattern, ".+")

		// 编译正则表达式
		re := regexp.MustCompile(pattern)

		// 检查输入是否匹配正则表达式
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

// IsBlacklist 判断token是否在黑名单
func (j *jwt) IsBlacklist(token string) bool {
	if j.redis == nil {
		return false
	}
	key := fmt.Sprintf("%s:%s", blackPrefix, token)
	is, _ := j.redis.Get(context.Background(), key).Result()
	return is == "1"
}

// AddBlacklist 添加token进入黑名单
func (j *jwt) AddBlacklist(token string) {
	key := fmt.Sprintf("%s:%s", blackPrefix, token)
	j.redis.Set(context.Background(), key, "1", j.conf.Expire)
}

// SetUnique 设置当前token为unique token
func (j *jwt) SetUnique(key, token string) error {
	km := crypto.MD5([]byte(key))
	tm := crypto.MD5([]byte(token))
	return j.redis.Set(context.Background(), uniquePrefix+":"+km, tm, j.conf.Expire).Err()
}

// CompareUniqueToken 对比是否时unique key
func (j *jwt) CompareUniqueToken(key, token string) bool {
	km := crypto.MD5([]byte(key))
	tm := crypto.MD5([]byte(token))
	res, _ := j.redis.HGet(context.Background(), uniquePrefix+":"+km, km).Result()
	return res == tm
}

// GetInfo 从ctx中获取token
func (j *jwt) GetInfo(ctx context.Context, dst any) error {
	if md, ok := metadata.FromServerContext(ctx); ok {
		body := md.Get(infoMdKey)
		if err := json.Unmarshal([]byte(body), dst); err != nil {
			return errors.New("auth info format error:" + err.Error())
		}
		return nil
	}
	return errors.New("not exist auth info")
}

// SetInfo 设置token的值到当前的ctx
func (j *jwt) SetInfo(req *http.Request, data string) {
	if data == "" {
		return
	}
	req.Header.Set(infoMdKey, data)
}
