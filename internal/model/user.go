package model

import (
	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/pkg/tree"
	"github.com/limes-cloud/manager/pkg/util"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx"
)

type User struct {
	BaseModel
	DepartmentID uint32      `json:"department_id"`
	RoleID       uint32      `json:"role_id"`
	Name         string      `json:"name"`
	Nickname     string      `json:"nickname"`
	Gender       string      `json:"gender"`
	Phone        string      `json:"phone"`
	Password     string      `json:"password"`
	Avatar       string      `json:"avatar"`
	Email        string      `json:"email"`
	Status       *bool       `json:"status"`
	Disabled     *string     `json:"disabled"`
	LastLogin    uint32      `json:"last_login"`
	Token        string      `json:"token"`
	Role         *Role       `json:"role"`
	Department   *Department `json:"department"`
}

// OneByID 通过id查询用户信息
func (u *User) OneByID(ctx kratosx.Context, id uint32) error {
	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	return db.First(u, id).Error
}

// OneByPhone 通过phone查询用户信息
func (u *User) OneByPhone(ctx kratosx.Context, phone string) error {
	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	return db.First(u, "phone=?", phone).Error
}

// OneByEmail 通过email查询用户信息
func (u *User) OneByEmail(ctx kratosx.Context, email string) error {
	db := ctx.DB().Model(u).Preload("Role").Preload("Department")
	return db.First(u, "email=?", email).Error
}

// Page 查询分页数据
func (u *User) Page(ctx kratosx.Context, options *PageOptions) ([]*User, uint32, error) {
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
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除用户信息
func (u *User) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}

// DataScopeTree 用户所管理的部门树
func (u *User) DataScopeTree(ctx kratosx.Context, id uint32) ([]tree.Tree, error) {
	// 操作者信息
	user := User{}
	if err := user.OneByID(ctx, id); err != nil {
		return nil, err
	}

	// 查询角色信息
	role := Role{}
	if err := role.OneByID(ctx, user.RoleID); err != nil {
		return nil, err
	}

	// 只存在当前部门的管理权限
	if role.DataScope == consts.DATA_SCOPE_CURRENT {
		return []tree.Tree{user.Department}, nil
	}

	var list []*Department
	var trees []tree.Tree
	dep := Department{}

	// 自定义部门权限
	if role.DataScope == consts.DATA_SCOPE_CUSTOM {
		ids := make([]uint32, 0)
		arr := strings.Split(*role.DepartmentIds, ",")
		for _, item := range arr {
			pid, _ := strconv.Atoi(item)
			ids = append(ids, uint32(pid))
		}
		list, _ = dep.All(ctx, func(db *gorm.DB) *gorm.DB {
			return db.Where("id in ?", ids)
		})
	} else {
		list, _ = dep.All(ctx, nil)
	}

	for _, item := range list {
		trees = append(trees, item)
	}

	// 自定义部门
	if role.DataScope == consts.DATA_SCOPE_CUSTOM {
		var resTree []tree.Tree
		roots := tree.FindRoots(trees)
		for _, root := range roots {
			resTree = append(resTree, tree.BuildTreeByID(trees, root))
		}
		return resTree, nil
	}

	// 当前部门级以下部门
	if role.DataScope == consts.DATA_SCOPE_CURRENT_DOWN {
		return []tree.Tree{tree.BuildTreeByID(trees, id)}, nil
	}

	// 下级部门
	if role.DataScope == consts.DATA_SCOPE_DOWN {
		resTree := tree.BuildTreeByID(trees, id)
		return resTree.ChildrenNode(), nil
	}

	// 所有部门
	if role.DataScope == consts.DATA_SCOPE_ALL {
		return []tree.Tree{tree.BuildTree(trees)}, nil
	}

	return []tree.Tree{}, nil
}

// DataScope 通过用户id获取用户所管理的部门id
func (u *User) DataScope(ctx kratosx.Context, id uint32) ([]uint32, error) {
	// 操作者信息
	user := User{}
	if err := user.OneByID(ctx, id); err != nil {
		return nil, err
	}

	// 查询角色信息
	role := Role{}
	if err := role.OneByID(ctx, user.RoleID); err != nil {
		return nil, err
	}

	// 根据条件取值
	switch role.DataScope {
	case consts.DATA_SCOPE_ALL: // 所有部门
		dep := Department{}
		ids := make([]uint32, 0)
		deps, err := dep.All(ctx, nil)
		if err != nil {
			return nil, err
		}
		for _, item := range deps {
			ids = append(ids, item.ID())
		}
		return ids, nil
	case consts.DATA_SCOPE_CURRENT: // 当前部门
		return []uint32{user.DepartmentID}, nil
	case consts.DATA_SCOPE_CURRENT_DOWN: // 当前部门以及下级部门
		dep := Department{}
		depTree, err := dep.TreeByID(ctx, user.DepartmentID)
		if err != nil {
			return nil, err
		}
		return tree.GetTreeID(depTree), nil
	case consts.DATA_SCOPE_DOWN: // 下级部门
		dep := Department{}
		depTree, err := dep.TreeByID(ctx, user.DepartmentID)
		if err != nil {
			return nil, err
		}
		ids := tree.GetTreeID(depTree)
		if len(ids) >= 2 {
			return ids[1:], nil
		}
		return []uint32{}, nil
	case consts.DATA_SCOPE_CUSTOM:
		ids := make([]uint32, 0)
		arr := strings.Split(*role.DepartmentIds, ",")
		for _, item := range arr {
			pid, _ := strconv.Atoi(item)
			ids = append(ids, uint32(pid))
		}
		return ids, nil
	}
	return []uint32{}, nil
}

// RoleScope 通过用户id获取用户所管理的角色ID
func (u *User) RoleScope(ctx kratosx.Context, id uint32) ([]uint32, error) {
	// 操作者信息
	user := User{}
	if err := user.OneByID(ctx, id); err != nil {
		return nil, err
	}

	// 查询角色信息
	role := Role{}
	roleTree, err := role.TreeByID(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}
	return tree.GetTreeID(roleTree), nil
}

func (u *User) HasUserScope(ctx kratosx.Context, mid, id uint32) bool {
	user := User{}

	depIds, err := u.DataScope(ctx, mid)
	if err != nil {
		return false
	}

	// 查询指定用户是否在部门内
	return ctx.DB().Model(&user).Where("department_id in ? and id=?", depIds, id).First(&user).RowsAffected != 0
}
