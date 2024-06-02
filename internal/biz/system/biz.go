package system

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetSystemSetting 获取系统设置
func (u *UseCase) GetSystemSetting(ctx kratosx.Context, req *GetSystemSettingRequest) *GetSystemSettingReply {
	setting := u.conf.Setting
	reply := GetSystemSettingReply{
		Debug:              setting.Debug,
		Title:              setting.Title,
		Desc:               setting.Desc,
		Copyright:          setting.Copyright,
		Logo:               setting.Logo,
		Watermark:          u.conf.Setting.Watermark,
		ChangePasswordType: u.conf.ChangePasswordType,
	}

	if len(req.DictionaryKeywords) != 0 {
		dictionaries, err := u.repo.GetDictionaryValues(ctx, req.DictionaryKeywords)
		if err != nil {
			ctx.Logger().Error("get dictionary error ", err.Error())
		}
		reply.Dictionaries = dictionaries
	}

	return &reply

}
