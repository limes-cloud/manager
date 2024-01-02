package model

import (
	"strings"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/pkg/util"
)

type User struct {
	types.BaseModel
	DepartmentID uint32      `json:"department_id" gorm:"not null;comment:部门id"`
	RoleID       uint32      `json:"role_id" gorm:"not null;comment:角色id"`
	Name         string      `json:"name" gorm:"not null;size:32;comment:名称"`
	Nickname     string      `json:"nickname" gorm:"not null;size:64;comment:昵称"`
	Gender       string      `json:"gender" gorm:"not null;size:12;comment:性别"`
	Phone        string      `json:"phone" gorm:"not null;size:11;comment:手机号码"`
	Password     string      `json:"password" gorm:"not null;size:256;comment:密码"`
	Avatar       string      `json:"avatar" gorm:"size:128;comment:头像"`
	Email        string      `json:"email" gorm:"not null;size:64;comment:邮箱"`
	Status       *bool       `json:"status" gorm:"default:false;comment:状态"`
	Disabled     *string     `json:"disabled" gorm:"size:64;comment:禁用原因"`
	LastLogin    int64       `json:"last_login" gorm:"comment:禁用原因"`
	Token        string      `json:"token" gorm:"size:512;comment:禁用原因"`
	Role         *Role       `json:"role"`
	Department   *Department `json:"department"`
}

// FindByID 通过id查询用户信息
func (u *User) FindByID(ctx kratosx.Context, id uint32) error {
	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	return db.First(u, id).Error
}

// FindByPhone 通过phFind查询用户信息
func (u *User) FindByPhone(ctx kratosx.Context, phFind string) error {
	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	return db.First(u, "phone=?", phFind).Error
}

// FindByEmail 通过email查询用户信息
func (u *User) FindByEmail(ctx kratosx.Context, email string) error {
	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	return db.First(u, "email=?", email).Error
}

// Page 查询分页数据
func (u *User) Page(ctx kratosx.Context, options *types.PageOptions) ([]*User, uint32, error) {
	var list []*User
	total := int64(0)

	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// Create 创建用户信息
func (u *User) Create(ctx kratosx.Context) error {
	u.Password = util.ParsePwd(u.Password)
	return ctx.DB().Model(u).Create(u).Error
}

func (u *User) UpdateLastLogin(ctx kratosx.Context, t uint32) error {
	return ctx.DB().Model(u).Where("id", u.ID).Update("last_login", t).Error
}

// Update 更新用户信息
func (u *User) Update(ctx kratosx.Context) error {
	if u.Password != "" {
		u.Password = util.ParsePwd(u.Password)
	}
	return ctx.DB().Updates(&u).Error
}

// DeleteByID 通过id删除用户信息
func (u *User) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}

// ManagerDepartmentIds 通过用户id获取用户所管理的部门id列表
func (u *User) ManagerDepartmentIds(ctx kratosx.Context, uid uint32) ([]uint32, error) {
	// 查询用户信息
	user := User{}
	if err := user.FindByID(ctx, uid); err != nil {
		return nil, err
	}

	// 查询用户当前角色信息
	role := Role{}
	if err := role.FindByID(ctx, user.RoleID); err != nil {
		return nil, err
	}

	// 根据条件取值
	switch role.DataScope {
	case consts.DataScopeAll: // 所有部门
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("id").Model(Department{}).Scan(&ids).Error; err != nil {
			return nil, err
		}
		return ids, nil
	case consts.DataScopeCurrent: // 当前部门
		return []uint32{user.DepartmentID}, nil
	case consts.DataScopeCurrentDown: // 当前部门以及下级部门
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("children").Model(DepartmentClosure{}).Where("parent=?", user.DepartmentID).Scan(&ids).Error; err != nil {
			return nil, err
		}
		return append(ids, user.DepartmentID), nil
	case consts.DataScopeDown: // 下级部门
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("children").Model(DepartmentClosure{}).Where("parent=?", user.DepartmentID).Scan(&ids).Error; err != nil {
			return nil, err
		}
		return ids, nil
	case consts.DataScopeCustom:
		ids := make([]uint32, 0)
		arr := strings.Split(*role.DepartmentIds, ",")
		for _, id := range arr {
			ids = append(ids, util.ToUint32(id))
		}
		return ids, nil
	}
	return []uint32{}, nil
}

// ManagerDepartment 通过用户id获取用户所管理的部门列表
func (u *User) ManagerDepartment(ctx kratosx.Context, uid uint32) ([]*Department, error) {
	// 查询用户信息
	user := User{}
	if err := user.FindByID(ctx, uid); err != nil {
		return nil, err
	}

	// 查询用户当前角色信息
	role := Role{}
	if err := role.FindByID(ctx, user.RoleID); err != nil {
		return nil, err
	}

	// 根据条件取值
	switch role.DataScope {
	case consts.DataScopeAll: // 所有部门
		dt := Department{}
		return dt.All(ctx, nil)
	case consts.DataScopeCurrent: // 当前部门
		return []*Department{user.Department}, nil
	case consts.DataScopeCurrentDown: // 当前部门以及下级部门
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("children").Model(DepartmentClosure{}).Where("parent=?", user.DepartmentID).Scan(&ids).Error; err != nil {
			return nil, err
		}
		ids = append(ids, user.DepartmentID)
		dt := Department{}
		return dt.All(ctx, func(db *gorm.DB) *gorm.DB {
			return db.Where("id in ?", ids)
		})
	case consts.DataScopeDown: // 下级部门
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("children").Model(DepartmentClosure{}).Where("parent=?", user.DepartmentID).Scan(&ids).Error; err != nil {
			return nil, err
		}

		dt := Department{}
		return dt.All(ctx, func(db *gorm.DB) *gorm.DB {
			return db.Where("id in ?", ids)
		})
	case consts.DataScopeCustom:
		ids := make([]uint32, 0)
		arr := strings.Split(*role.DepartmentIds, ",")
		for _, id := range arr {
			ids = append(ids, util.ToUint32(id))
		}
		dt := Department{}
		return dt.All(ctx, func(db *gorm.DB) *gorm.DB {
			return db.Where("id in ?", ids)
		})
	}
	return []*Department{}, nil
}

// HasUserManagerScope 判断用户是否具有指定用户的管理权限
func (u *User) HasUserManagerScope(ctx kratosx.Context, mid, id uint32) bool {
	user := User{}

	depIds, err := u.ManagerDepartmentIds(ctx, mid)
	if err != nil {
		return false
	}

	// 查询指定用户是否在部门内
	return ctx.DB().Model(&user).Where("department_id in ? and id=?", depIds, id).First(&user).RowsAffected != 0
}

// HasRoleManagerScope 判断用户是否具有指定用户的管理权限
func (u *User) HasRoleManagerScope(ctx kratosx.Context, rid uint32) bool {
	role := &Role{
		BaseModel: types.BaseModel{ID: u.RoleID},
	}

	ids, _ := role.FindManagerIds(ctx)
	return util.InList(ids, rid)
}
