package user

import (
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/manager/internal/biz/department"
	"github.com/limes-cloud/manager/internal/biz/job"
	"github.com/limes-cloud/manager/internal/biz/role"
)

type User struct {
	types.BaseModel
	DepartmentId uint32                 `json:"department_id"`
	RoleId       uint32                 `json:"role_id"`
	Name         string                 `json:"name"`
	Nickname     string                 `json:"nickname"`
	Gender       string                 `json:"gender"`
	Phone        string                 `json:"phone"`
	Password     string                 `json:"password"`
	Avatar       string                 `json:"avatar"`
	Email        string                 `json:"email"`
	Status       *bool                  `json:"status"`
	Disabled     *string                `json:"disabled"`
	LastLogin    int64                  `json:"last_login"`
	Token        string                 `json:"token"`
	Role         *role.Role             `json:"role"`
	Department   *department.Department `json:"department"`
	Roles        []*role.Role           `json:"roles" gorm:"many2many:user_role"`
	Jobs         []*job.Job             `json:"jobs" gorm:"many2many:user_job"`
	UserRoles    []*UserRole            `json:"user_roles"`
	UserJobs     []*UserJob             `json:"user_jobs"`
}

type UserRole struct {
	types.CreateModel
	RoleID uint32 `json:"role_id"`
	UserID uint32 `json:"user_id"`
}

type UserJob struct {
	types.CreateModel
	JobID  uint32 `json:"job_id"`
	UserID uint32 `json:"user_id"`
}
