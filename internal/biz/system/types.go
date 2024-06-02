package system

type GetSystemSettingRequest struct {
	DictionaryKeywords []string `json:"dictionaryKeywords"`
}

type GetSystemSettingReply struct {
	Debug              bool               `json:"debug"`
	Title              string             `json:"title"`
	Desc               string             `json:"desc"`
	Copyright          string             `json:"copyright"`
	Logo               string             `json:"logo"`
	ChangePasswordType string             `json:"changePasswordType"`
	Watermark          string             `json:"watermark"`
	Dictionaries       []*DictionaryValue `json:"dictionaries"`
}

type DictionaryValue struct {
	Keyword string  `json:"keyword"`
	Label   string  `json:"label"`
	Value   string  `json:"value"`
	Type    *string `json:"type"`
	Extra   *string `json:"extra"`
}
