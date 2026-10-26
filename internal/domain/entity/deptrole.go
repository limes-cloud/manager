package entity

type DeptRole struct {
	ID     uint32 `json:"id" gorm:"column:id"`
	App    string `json:"app" gorm:"column:app"`
	DeptId uint32 `json:"deptId" gorm:"column:dept_id"`
	RoleId uint32 `json:"roleId"  gorm:"column:role_id"`
}
