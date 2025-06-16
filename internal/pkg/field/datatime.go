package field

import (
	"time"

	"google.golang.org/protobuf/types/known/structpb"
)

type datetimeType struct {
}

func (dt datetimeType) Name() string {
	return "日期时间"
}

func (dt datetimeType) Validate(in *structpb.Value) bool {
	_, err := time.Parse("2006-01-02 15:04:05", in.GetStringValue())
	return err == nil
}

func (dt datetimeType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (dt datetimeType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
