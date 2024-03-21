package user

type PageUserRequest struct {
	Page          uint32   `json:"page"`
	PageSize      uint32   `json:"page_size"`
	DepartmentIds []uint32 `json:"department_ids"`
	Username      *string  `json:"username"`
	Status        *bool    `json:"status"`
	RoleId        *uint32  `json:"role_id"`
	DepartmentId  *uint32  `json:"department_id"`
	StartTime     *uint32  `json:"start_time"`
	EndTime       *uint32  `json:"end_time"`
	Name          *string  `json:"name"`
}

type UpdateCurrentUserRequest struct {
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
}

type CaptchaReply struct {
	Uuid    string `json:"uuid"`
	Captcha string `json:"captcha"`
	Expire  uint32 `json:"expire"`
}

type ChangeUserPasswordRequest struct {
	Password  string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

type UserLoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}
