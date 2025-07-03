package entity

import "github.com/limes-cloud/kratosx/types"

type User struct {
	Name            string            `json:"name" gorm:"column:name"`
	Nickname        string            `json:"nickname" gorm:"column:nickname"`
	Gender          string            `json:"gender" gorm:"column:gender"`
	Avatar          *string           `json:"avatar" gorm:"column:avatar"`
	Phone           string            `json:"phone" gorm:"column:phone"`
	Email           string            `json:"email" gorm:"column:email"`
	Password        string            `json:"password" gorm:"column:password"`
	Status          *bool             `json:"status" gorm:"column:status"`
	Setting         *string           `json:"setting" gorm:"column:setting"`
	Token           *string           `json:"token" gorm:"column:token"`
	LoggedAt        int64             `json:"loggedAt" gorm:"column:logged_at"`
	ExpireAt        int64             `json:"expireAt" gorm:"column:expire_at"`
	UserJobs        []*UserJob        `json:"userJobs"`
	UserRoles       []*UserRole       `json:"userRoles"`
	UserDepartments []*UserDepartment `json:"userDepartments"`
	Roles           []*Role           `json:"roles" gorm:"many2many:user_role"`
	Jobs            []*Job            `json:"jobs" gorm:"many2many:user_job"`
	Departments     []*Department     `json:"departments" gorm:"many2many:user_department"`
	types.BaseModel
}

func (u User) GetRoleIds() []uint32 {
	var ids []uint32
	if len(u.UserRoles) != 0 {
		for _, item := range u.UserRoles {
			ids = append(ids, item.RoleId)
		}
		return ids
	}
	for _, item := range u.Roles {
		ids = append(ids, item.Id)
	}
	return ids
}

func (u User) GetJobIds() []uint32 {
	var ids []uint32
	if len(u.UserJobs) != 0 {
		for _, item := range u.UserJobs {
			ids = append(ids, item.JobId)
		}
		return ids
	}
	for _, item := range u.Jobs {
		ids = append(ids, item.Id)
	}
	return ids
}

func (u User) GetDepartmentIds() []uint32 {
	var ids []uint32

	if len(u.UserDepartments) != 0 {
		for _, item := range u.UserDepartments {
			ids = append(ids, item.DepartmentId)
		}
		return ids
	}

	for _, item := range u.Departments {
		ids = append(ids, item.Id)
	}
	return ids
}

type UserDepartment struct {
	UserId       uint32 `json:"userId" gorm:"column:user_id"`
	DepartmentId uint32 `json:"departmentId" gorm:"column:department_id"`
}

type UserJob struct {
	UserId uint32 `json:"userId" gorm:"column:user_id"`
	JobId  uint32 `json:"jobId" gorm:"column:job_id"`
}

type UserRole struct {
	UserId uint32 `json:"userId" gorm:"column:user_id"`
	RoleId uint32 `json:"roleId" gorm:"column:role_id"`
}
