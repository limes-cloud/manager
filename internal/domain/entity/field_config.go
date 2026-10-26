package entity

// FieldConfig 组件配置（统一存放所有配置项）
type FieldConfig struct {
	Placeholder  string        `json:"placeholder,omitempty"`  // 提示文本
	DefaultValue any           `json:"defaultValue,omitempty"` // 默认值（string/number/array/null）
	MaxLength    *uint32       `json:"maxLength,omitempty"`    // input/textarea/password: 最大长度
	Min          *float64      `json:"min,omitempty"`          // number: 最小值
	Max          *float64      `json:"max,omitempty"`          // number: 最大值
	Step         *float64      `json:"step,omitempty"`         // number: 步长
	Options      []FieldOption `json:"options,omitempty"`      // radio/checkbox/select: 选项列表
	Filetype     string        `json:"filetype,omitempty"`     // upload: 文件类型（image/video/file）
	Limit        *uint32       `json:"limit,omitempty"`        // upload: 最大文件数
	Rule         *FieldRule    `json:"rule,omitempty"`         // 校验规则
}

// FieldOption 选项（radio/checkbox/select）
type FieldOption struct {
	Label string `json:"label"` // 显示文本
	Value string `json:"value"` // 值
}

// FieldRule 校验规则
type FieldRule struct {
	Type    string   `json:"type"`              // phone/email/idcard/regex/enum/range/length
	Message string   `json:"message,omitempty"` // 自定义错误提示
	Pattern string   `json:"pattern,omitempty"` // type=regex 时的正则表达式
	Values  []string `json:"values,omitempty"`  // type=enum 时的合法值列表
	Min     *float64 `json:"min,omitempty"`     // type=range/length 时的最小值
	Max     *float64 `json:"max,omitempty"`     // type=range/length 时的最大值
}
