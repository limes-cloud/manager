package util

import (
	"testing"
)

func TestParsePwd(t *testing.T) {
	res := ParsePwd("12345678")
	t.Log(res, len(res))
}

// func TestCompareValue(t *testing.T) {
//	tp := new(string)
//	CompareValue(tp, "a")
// }
