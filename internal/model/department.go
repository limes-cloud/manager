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

type Department struct {
	types.BaseModel
	ParentID    uint32        `json:"parent_id" gorm:"not null;comment:父id"`
	Path        string        `json:"path" gorm:"not null;size:256;comment:层级路径"`
	Keyword     string        `json:"keyword" gorm:"not null;size:32;unique;comment:关键字"`
	Name        string        `json:"name" gorm:"not null;size:32;unique;comment:名称"`
	Description string        `json:"description" gorm:"not null;size:256;comment:描述"`
	Children    []*Department `json:"children"  gorm:"-"`
}

type DepartmentClosure struct {
	ID                 uint32      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Parent             uint32      `json:"parent" gorm:"not null;comment:部门id"`
	Children           uint32      `json:"children" gorm:"not null;comment:部门id"`
	ParentDepartment   *Department `json:"parent_department" gorm:"foreignKey:parent;references:id"`
	ChildrenDepartment *Department `json:"children_department" gorm:"foreignKey:children;references:id"`
}

func (t *Department) ID() uint32 {
	return t.BaseModel.ID
}

func (t *Department) Parent() uint32 {
	return t.ParentID
}

func (t *Department) AppendChildren(child any) {
	team := child.(*Department)
	t.Children = append(t.Children, team)
}

func (t *Department) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range t.Children {
		list = append(list, item)
	}
	return list
}

// All 获取全部部门
func (t *Department) All(ctx kratosx.Context, scopes types.Scopes) ([]*Department, error) {
	list := make([]*Department, 0)
	db := ctx.DB().Model(Department{})
	if scopes != nil {
		db = scopes(db)
	}
	return list, db.Find(&list).Error
}

// Create 创建部门
func (t *Department) Create(ctx kratosx.Context) error {
	if t.ParentID != 0 {
		pt := Department{}
		if err := pt.FindByID(ctx, t.ParentID); err != nil {
			return errors.New("父菜单不存在")
		}
		t.Path = fmt.Sprintf("%s%d/", pt.Path, pt.ID())
	} else {
		t.Path = "/"
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(t).Error; err != nil {
			return err
		}

		dcs := t.departmentClosure(t.Path, t.ID())
		if len(dcs) != 0 {
			return tx.Create(&dcs).Error
		}
		return nil
	})
}

// FindByID 通过部门id查询
func (t *Department) FindByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(t, "id=?", id).Error
}

// FindByKeyword 通过部门keyword查询
func (t *Department) FindByKeyword(ctx kratosx.Context, keyword string) error {
	return ctx.DB().First(t, "keyword=?", keyword).Error
}

// Update 更新部门信息
func (t *Department) Update(ctx kratosx.Context) error {
	old := Department{}
	if err := old.FindByID(ctx, t.ID()); err != nil {
		return err
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		// 存在层级变更
		if old.ParentID != t.ParentID {
			// 将原来的上层级取消
			if err := tx.Delete(DepartmentClosure{}, "children=?", t.ID()).Error; err != nil {
				return err
			}

			// 查询当前的父节点
			parent := Department{}
			if err := parent.FindByID(ctx, t.ParentID); err != nil {
				return err
			}

			// 重置path
			t.Path = fmt.Sprintf("%s%d/", parent.Path, parent.ID())

			// 新增现在的层级
			dcs := t.departmentClosure(t.Path, t.ID())
			if len(dcs) != 0 {
				return tx.Create(&dcs).Error
			}
		}
		return tx.Updates(t).Error
	})
}

// DeleteByID 通过id删除指定部门 以及部门下的部门
func (t *Department) DeleteByID(ctx kratosx.Context, id uint32) error {
	// 获取下级id
	ids := make([]uint32, 0)
	if err := ctx.DB().Select("children").Model(DepartmentClosure{}).Where("parent=?", id).Scan(&ids).Error; err != nil {
		return err
	}
	ids = append(ids, id)

	return ctx.DB().Where("id in ?", ids).Delete(Department{}).Error
}

func (t *Department) getParentIdsByPath(path string) []uint32 {
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

func (t *Department) departmentClosure(path string, id uint32) []DepartmentClosure {
	var dcs []DepartmentClosure

	arr := t.getParentIdsByPath(path)
	for _, pid := range arr {
		dcs = append(dcs, DepartmentClosure{
			Parent:   pid,
			Children: id,
		})
	}

	return dcs
}
