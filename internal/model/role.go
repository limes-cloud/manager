package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/pkg/tree"
	"github.com/limes-cloud/manager/pkg/util"
)

type Role struct {
	types.BaseModel
	ParentID      uint32  `json:"parent_id" gorm:"not null;comment:父id"`
	Name          string  `json:"name" gorm:"not null;type:char(64);comment:名称"`
	Keyword       string  `json:"keyword" gorm:"not null;type:char(32);comment:关键字"`
	Path          string  `json:"path" gorm:"not null;size:256;comment:层级路径"`
	Status        *bool   `json:"status" gorm:"not null;default:false;comment:状态"`
	Description   *string `json:"description" gorm:"not null;size:128;comment:描述"`
	DepartmentIds *string `json:"department_ids" gorm:"size:256;comment:自定义管理部门"`
	DataScope     string  `json:"data_scope"  gorm:"not null;type:char(32);comment:权限类型"`
	Children      []*Role `json:"children" gorm:"-"`
}

type RoleClosure struct {
	ID           uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Parent       uint32 `json:"parent" gorm:"not null;comment:用户id"`
	Children     uint32 `json:"children" gorm:"not null;comment:用户id"`
	ParentRole   *Role  `json:"parent_role" gorm:"->;foreignKey:parent;references:id"`
	ChildrenRole *Role  `json:"children_role" gorm:"->;foreignKey:children;references:id"`
}

// ID 获取ID
func (r *Role) ID() uint32 {
	return r.BaseModel.ID
}

// Parent 获取父ID
func (r *Role) Parent() uint32 {
	return r.ParentID
}

// AppendChildren 添加子树
func (r *Role) AppendChildren(child any) {
	menu := child.(*Role)
	r.Children = append(r.Children, menu)
}

// ChildrenNode 获取子树列表
func (r *Role) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range r.Children {
		list = append(list, item)
	}
	return list
}

// Create 创建角色信息
func (r *Role) Create(ctx kratosx.Context) error {
	if r.ParentID != 0 {
		pt := Department{}
		if err := pt.FindByID(ctx, r.ParentID); err != nil {
			return errors.New("父菜单不存在")
		}
		r.Path = fmt.Sprintf("%s%d/", pt.Path, pt.ID())
	} else {
		r.Path = "/"
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(r).Error; err != nil {
			return err
		}

		dcs := r.roleClosure(r.Path, r.ID())
		if len(dcs) != 0 {
			return tx.Create(&dcs).Error
		}
		return nil
	})
}

// FindManagerIds 查询用户管理的id列表
func (r *Role) FindManagerIds(ctx kratosx.Context) ([]uint32, error) {
	var ids []uint32
	if err := ctx.DB().Model(RoleClosure{}).
		Select("children").
		Where("parent=?", r.ID()).
		Scan(&ids).Error; err != nil {
		return ids, err
	}
	ids = append(ids, r.ID())
	return ids, nil
}

// FindByID 通过ID查询角色信息
func (r *Role) FindByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(r, "id=?", id).Error
}

// All 查询全部角色信息
func (r *Role) All(ctx kratosx.Context, scopes types.Scopes) ([]*Role, error) {
	var list []*Role
	db := ctx.DB()
	if scopes != nil {
		db.Scopes(scopes)
	}
	return list, db.Find(&list).Error
}

// Update 更新角色信息
func (r *Role) Update(ctx kratosx.Context) error {
	old := Role{}
	if err := old.FindByID(ctx, r.ID()); err != nil {
		return err
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		// 存在层级变更
		if old.ParentID != r.ParentID {
			// 将原来的上层级取消
			if err := tx.Delete(DepartmentClosure{}, "children=?", r.ID()).Error; err != nil {
				return err
			}

			// 查询当前的父节点
			parent := Department{}
			if err := parent.FindByID(ctx, r.ParentID); err != nil {
				return err
			}

			// 重置path
			r.Path = fmt.Sprintf("%s%d/", parent.Path, parent.ID())

			// 新增现在的层级
			dcs := r.roleClosure(r.Path, r.ID())
			if len(dcs) != 0 {
				return tx.Create(&dcs).Error
			}
		}
		return tx.Updates(r).Error
	})
}

// DeleteByID 通过ID删除角色信息
func (r *Role) DeleteByID(ctx kratosx.Context, id uint32) error {
	ids := make([]uint32, 0)
	if err := ctx.DB().Select("children").Model(RoleClosure{}).Where("parent=?", id).Scan(&ids).Error; err != nil {
		return err
	}
	ids = append(ids, id)

	return ctx.DB().Where("id in ?", ids).Delete(Role{}).Error
}

// ParentStatus 获取指定角色的父角色状态
func (r *Role) ParentStatus(ctx kratosx.Context) bool {
	ids := r.getParentIdsByPath(r.Path)
	total := int64(0)
	ctx.DB().Where("id in ? and status=false", ids).Count(&total)
	return total == 0
}

func (r *Role) getParentIdsByPath(path string) []uint32 {
	var pid []uint32
	if path == "/" {
		return pid
	}
	arr := strings.Split(path[1:len(path)-1], "/")
	for _, id := range arr {
		pid = append(pid, util.ToUint32(id))
	}
	return pid
}

func (r *Role) roleClosure(path string, id uint32) []RoleClosure {
	var dcs []RoleClosure

	arr := r.getParentIdsByPath(path)
	for _, pid := range arr {
		dcs = append(dcs, RoleClosure{
			Parent:   pid,
			Children: id,
		})
	}

	return dcs
}
