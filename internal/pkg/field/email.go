package field

import (
	"regexp"

	"google.golang.org/protobuf/types/known/structpb"
)

type emailType struct {
}

func (et emailType) Name() string {
	return "邮箱"
}

func (et emailType) Validate(in *structpb.Value) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(.[a-zA-Z]{2,})+$`)
	return reg.MatchString(in.GetStringValue())
}

func (et emailType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (et emailType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
