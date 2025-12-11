package entity

import (
	"fmt"

	"github.com/limes-cloud/kratosx/model"
)

type User struct {
	Avatar    string      `json:"avatar" gorm:"column:avatar"`                       // 用户头像
	Nickname  string      `json:"nickname" gorm:"column:nickname"`                   // 用户昵称
	Username  string      `json:"username" gorm:"column:username"`                   // 用户账户
	Password  string      `json:"password" gorm:"column:password"`                   // 用户密码
	Status    *bool       `json:"status" gorm:"column:status"`                       // 用户状态
	Reason    *string     `json:"reason" gorm:"column:reason"`                       // 用户原因
	LoggedAt  int64       `json:"loggedAt" gorm:"column:logged_at"`                  // 最近登录时间
	Setting   string      `json:"setting" gorm:"-"`                                  // 用户设置
	UserDepts []*UserDept `json:"userDepts" gorm:"foreignKey:user_id;references:id"` // 用户部门
	Infos     []*Userinfo `json:"infos" gorm:"-"`                                    // 用户信息
	Authorize *Authorize  `json:"authorize" gorm:"foreignKey:user_id;references:id"` // 用户授权
	model.BaseTenantModel
}

type UserSetting struct {
	UserId  uint32 `json:"userId" gorm:"column:user_id"`
	AppId   uint32 `json:"appId" gorm:"column:app_id"`
	Setting string `json:"setting" gorm:"column:setting"`
	model.BaseModel
}

type Userinfo struct {
	TenantId uint32 `json:"tenantId" gorm:"column:tenant_id"`
	UserId   uint32 `json:"userId" gorm:"column:user_id"`
	Field    string `json:"field" gorm:"column:field"`
	Value    string `json:"value" gorm:"column:value"`
	ValueMd5 string `json:"valueMd5" gorm:"column:value_md5"`
	Index    *int   `json:"index" gorm:"-"`
	model.BaseModel
}

func (u *Userinfo) GetTableName() string {
	if u.UserId != 0 && u.Index == nil {
		index := int(u.UserId % 10)
		u.Index = &index
	}
	return fmt.Sprintf("userinfo_%d", *u.Index)
}
