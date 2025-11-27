package entity

import (
	"encoding/json"
	"time"

	"github.com/limes-cloud/kratosx/model"
)

type Authorize struct {
	model.CreateTenantModel
	AppId     uint32 `json:"appId" gorm:"column:app_id"`
	UserId    uint32 `json:"userId" gorm:"column:user_id"`
	Tokens    string `json:"tokens" gorm:"column:tokens"`
	LoggedAt  int64  `json:"loggedAt" gorm:"column:logged_at"`
	ExpiredAt int64  `json:"expiredAt" gorm:"column:expired_at"`
}

func (a *Authorize) GetTokens() map[string]int64 {
	tm := make(map[string]int64)
	_ = json.Unmarshal([]byte(a.Tokens), &tm)

	// 清理已经超时的
	for k, v := range tm {
		if v < time.Now().Unix() {
			delete(tm, k)
		}
	}

	return tm
}
