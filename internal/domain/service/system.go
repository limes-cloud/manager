package service

import (
	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/types"
)

type System struct {
	// dict repository.Dictionary
}

func NewSystem() *System {
	return &System{}
}

// GetSystemSetting 获取系统设置
func (u *System) GetSystemSetting(ctx core.Context) *types.GetSystemSettingReply {
	setting := ctx.Config().Setting
	reply := types.GetSystemSettingReply{
		Debug:       setting.Debug,
		Title:       setting.Title,
		Description: setting.Desc,
		Copyright:   setting.Copyright,
		Logo:        setting.Logo,
		Watermark:   setting.Watermark,
		// ChangePasswordType: u.conf.ChangePasswordType,
	}

	//if len(u.conf.DictionaryKeywords) != 0 {
	//	dictionaries, err := u.dict.ListDictionaryValues(ctx, u.conf.DictionaryKeywords)
	//	if err != nil {
	//		ctx.Logger().Error("get dictionary error ", err.Error())
	//	}
	//	reply.Dictionaries = dictionaries
	//}

	return &reply
}
