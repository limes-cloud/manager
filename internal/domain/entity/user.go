package entity

import "github.com/limes-cloud/kratosx/model"

const (
	UserEntityName = "user"
)

type User struct {
	Avatar   *string `json:"avatar" gorm:"column:avatar"`      // 用户头像
	Nickname string  `json:"nickname" gorm:"column:nickname"`  // 用户昵称
	Username string  `json:"username" gorm:"column:username"`  // 用户账户
	Password string  `json:"password" gorm:"column:password"`  // 用户密码
	Status   *bool   `json:"status" gorm:"column:status"`      // 用户状态
	Token    *string `json:"token" gorm:"column:token"`        // 认证令牌
	LoggedAt int64   `json:"loggedAt" gorm:"column:logged_at"` // 登录时间
	ExpireAt int64   `json:"expireAt" gorm:"column:expire_at"` // 过期时间
	DeptId   uint32  `json:"deptId" gorm:"column:dept_id"`     // 所属部门
	JobId    uint32  `json:"jobId" gorm:"column:job_id"`       // 所属岗位
	Dept     *Dept   `json:"dept" gorm:"foreignKey:dept_id;references:id"`
	Job      *Job    `json:"job" gorm:"foreignKey:job_id;references:id"`
	model.BaseTenantDeptModel
}

type UserDept struct {
	UserId uint32 `json:"user_id" gorm:"column:user_id"`
	DeptId uint32 `json:"dept_id" gorm:"column:dept_id"`
	JobId  uint32 `json:"job_id" gorm:"column:job_id"`
	model.BaseModel
}

// TableName 获取表名称
func (m *User) TableName() string {
	return UserEntityName
}
