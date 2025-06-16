package field

import (
	"regexp"
	"strconv"

	"google.golang.org/protobuf/types/known/structpb"
)

type numberType struct {
}

func (nt numberType) Name() string {
	return "数字"
}

func (nt numberType) Validate(in *structpb.Value) bool {
	reg := regexp.MustCompile(`^-?\d*\.?\d+$`)
	return reg.MatchString(in.GetStringValue())
}

func (nt numberType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (nt numberType) ToValue(in string) *structpb.Value {
	value, _ := strconv.ParseFloat(in, 64)
	return structpb.NewNumberValue(value)
}
