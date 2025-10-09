package field

import (
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

type genderType struct{}

func (g genderType) Name() string {
	return "性别"
}

func (g genderType) Validate(in *structpb.Value) bool {
	str := in.GetStringValue()
	return str == "M" || str == "F"
}

func (g genderType) ToString(in *structpb.Value) string {
	return fmt.Sprint(in.GetStringValue())
}

func (g genderType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
