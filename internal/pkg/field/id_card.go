package field

import (
	"regexp"

	"google.golang.org/protobuf/types/known/structpb"
)

type idCardType struct {
}

func (it idCardType) Name() string {
	return "身份证"
}

func (it idCardType) Validate(in *structpb.Value) bool {
	reg := regexp.MustCompile(`^([1-9]\d{5})(\d{4})(0[1-9]|1[0-2])([0-2][1-9]|[1-3]\d|4[0-6]|5[0-2])(\d{3})(\d|[Xx])$`)
	return reg.MatchString(in.GetStringValue())
}

func (it idCardType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (it idCardType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
