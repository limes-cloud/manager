package user

type User struct {
	Id           uint32      `json:"id"`
	DepartmentId uint32      `json:"departmentId"`
	RoleId       uint32      `json:"roleId"`
	Name         string      `json:"name"`
	Nickname     string      `json:"nickname"`
	Gender       string      `json:"gender"`
	Avatar       *string     `json:"avatar"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	Status       *bool       `json:"status"`
	Setting      *string     `json:"setting"`
	Token        *string     `json:"token"`
	LoggedAt     int64       `json:"loggedAt"`
	CreatedAt    int64       `json:"createdAt"`
	UpdatedAt    int64       `json:"updatedAt"`
	UserJobs     []*UserJob  `json:"userJobs"`
	UserRoles    []*UserRole `json:"userRoles"`
	Roles        []*Role     `json:"roles"`
	Jobs         []*Job      `json:"jobs"`
	Department   *Department `json:"department"`
	Role         *Role       `json:"role"`
}

type Department struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
}

type Job struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
}

type Role struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
	Status  *bool  `json:"status"`
}

type UserJob struct {
	JobId uint32 `json:"jobId"`
}

type UserRole struct {
	RoleId uint32 `json:"roleId"`
}
