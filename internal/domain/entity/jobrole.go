package entity

type JobRole struct {
	ID     uint32 `json:"id" gorm:"column:id"`
	JobId  uint32 `json:"jobId" gorm:"column:job_id"`
	RoleId uint32 `json:"roleId"  gorm:"column:role_id"`
}
