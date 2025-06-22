package oauth

import (
	"crypto/md5"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type captcha struct {
}

// randomCode 生成随机数验证码
func (c *captcha) randomCode(len int) string {
	var code = rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

// sid 生成当前场景下客户端的唯一id
func (c *captcha) sid(scene, ip string) string {
	return fmt.Sprintf("captcha:s:%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", scene, ip))))
}

// sid 生成当前场景下客户端手机号的唯一id
func (c *captcha) uid(sid, user string) string {
	return fmt.Sprintf("captcha:u:%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", sid, user))))
}

// sid 生成当前场景下客户端手机号的唯一id
func (c *captcha) getUUIDByUID(uid string) string {
	return strings.TrimPrefix(uid, "captcha:u:")
}

// getUidByUuid 生成验证码的唯一uid
func (c *captcha) getUIDByUUID(uuid string) string {
	return fmt.Sprintf("captcha:u:%s", uuid)
}

// uuid 生成验证码的唯一uid
func (c *captcha) aid(scene, user, answer string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", scene, user, answer))))
}
