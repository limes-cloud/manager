package types

type GetSystemSettingRequest struct{}

type GetSystemSettingReply struct {
	Debug              bool   `json:"debug"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Copyright          string `json:"copyright"`
	Logo               string `json:"logo"`
	ChangePasswordType string `json:"changePasswordType"`
	Watermark          string `json:"watermark"`
	// Dictionaries       []*DictionaryValue `json:"dictionaries"`
}
