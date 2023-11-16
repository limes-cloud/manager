package util

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"

	json "github.com/json-iterator/go"
)

func Transform(in interface{}, dst interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func MD5(in []byte) string {
	sum := md5.Sum(in)
	return hex.EncodeToString(sum[:])
}

func MD5ToUpper(in []byte) string {
	return strings.ToUpper(MD5(in))
}

func ParsePwd(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

type ListType interface {
	~string | ~int | ~uint32 | ~[]byte | ~rune | ~float64
}

func InList[ListType comparable](list []ListType, val ListType) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func IsEmail(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(.[a-zA-Z]{2,})+$`)
	return reg.MatchString(email)
}

func IsPhone(phone string) bool {
	reg := regexp.MustCompile(`^1[3456789]\d{9}$`)
	return reg.MatchString(phone)
}

//type CompareType interface {
//	~string | ~int | ~uint32 | ~float64 | ~*string | ~*int | ~*uint32 | ~*float64
//}
//
//func CompareValue[CompareType comparable](v1 CompareType, v2 CompareType) bool {
//	return v1 == v2
//}