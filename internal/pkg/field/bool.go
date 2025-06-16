package field

import (
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

type boolType struct {
}

func (boolType boolType) Name() string {
	return "布尔"
}

func (boolType boolType) Validate(in *structpb.Value) bool {
	str := in.GetStringValue()
	return str == "true" || str == "false"
}

func (boolType boolType) ToString(in *structpb.Value) string {
	return fmt.Sprint(in.GetBoolValue())
}

func (boolType boolType) ToValue(in string) *structpb.Value {
	return structpb.NewBoolValue(in == "true")
}
