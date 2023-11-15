package util

import (
	"testing"
)

func TestParsePwd(t *testing.T) {
	res := ParsePwd("123456")
	t.Log(res, len(res))
}

//func TestCompareValue(t *testing.T) {
//	tp := new(string)
//	CompareValue(tp, "a")
//}
