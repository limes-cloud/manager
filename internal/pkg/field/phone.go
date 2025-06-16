package field

import (
	"regexp"

	"google.golang.org/protobuf/types/known/structpb"
)

type phoneType struct {
}

func (pt phoneType) Name() string {
	return "手机号"
}

func (pt phoneType) Validate(in *structpb.Value) bool {
	reg := regexp.MustCompile(`^1[3456789]\d{9}$`)
	return reg.MatchString(in.GetStringValue())
}

func (pt phoneType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (pt phoneType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
