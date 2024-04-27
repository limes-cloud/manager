package department

import (
	"strings"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/manager/internal/biz/department"
	userBiz "github.com/limes-cloud/manager/internal/biz/user"
)

func NewRepo() biz.Repo {
	return &repo{}
}

type repo struct {
}

func (r repo) IsManagerAllDepartment(ctx kratosx.Context, uid uint32) (bool, error) {
	user := userBiz.User{}
	if err := ctx.DB().Preload("Role").First(&user, uid).Error; err != nil {
		return false, err
	}
	return user.Role.DataScope == biz.DataScopeAll || user.DepartmentId == 1, nil
}

func (r repo) AllManagerDepartmentIds(ctx kratosx.Context, uid uint32) ([]uint32, error) {
	// 查询用户信息
	user := userBiz.User{}
	if err := ctx.DB().Preload("Role").Preload("Department").First(&user, uid).Error; err != nil {
		return nil, err
	}

	// 根据条件取值
	switch user.Role.DataScope {
	case biz.DataScopeAll: // 所有部门
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("id").Model(biz.Department{}).Scan(&ids).Error; err != nil {
			return nil, err
		}
		return ids, nil
	case biz.DataScopeCurrent: // 当前部门
		return []uint32{user.DepartmentId}, nil
	case biz.DataScopeCurrentDown: // 当前部门以及下级部门
		return r.GetChildrenIds(ctx, user.DepartmentId)
	case biz.DataScopeDown: // 下级部门
		ids, err := r.GetChildrenIds(ctx, user.DepartmentId)
		if err != nil {
			return ids, err
		}
		if len(ids) >= 1 {
			ids = ids[:len(ids)-1]
		}
		return ids, nil
	case biz.DataScopeCustom:
		ids := make([]uint32, 0)
		arr := strings.Split(*user.Role.DepartmentIds, ",")
		for _, id := range arr {
			ids = append(ids, util.ToUint32(id))
		}
		return ids, nil
	}
	return []uint32{}, nil
}

func (r repo) AllManagerDepartment(ctx kratosx.Context, uid uint32) ([]*biz.Department, error) {
	// 查询用户信息
	user := userBiz.User{}
	if err := ctx.DB().Preload("Role").Preload("Department").First(&user, uid).Error; err != nil {
		return nil, err
	}

	// 根据条件取值
	switch user.Role.DataScope {
	case biz.DataScopeAll: // 所有部门
		return r.AllDepartment(ctx, nil)
	case biz.DataScopeCurrent: // 当前部门
		return []*biz.Department{user.Department}, nil
	case biz.DataScopeCurrentDown: // 当前部门以及下级部门
		ids, err := r.GetChildrenIds(ctx, user.DepartmentId)
		if err != nil {
			return nil, err
		}
		return r.AllDepartment(ctx, &biz.AllDepartmentRequest{Ids: ids})
	case biz.DataScopeDown: // 下级部门
		ids, err := r.GetChildrenIds(ctx, user.DepartmentId)
		if err != nil {
			return nil, err
		}
		if len(ids) >= 1 {
			ids = ids[:len(ids)-1]
		}
		return r.AllDepartment(ctx, &biz.AllDepartmentRequest{Ids: ids})
	case biz.DataScopeCustom:
		ids := make([]uint32, 0)
		arr := strings.Split(*user.Role.DepartmentIds, ",")
		for _, id := range arr {
			ids = append(ids, util.ToUint32(id))
		}
		return r.AllDepartment(ctx, &biz.AllDepartmentRequest{Ids: ids})
	}
	return []*biz.Department{}, nil
}

func (r repo) GetDepartment(ctx kratosx.Context, id uint32) (*biz.Department, error) {
	var department biz.Department
	return &department, ctx.DB().First(&department, id).Error
}

func (r repo) AllDepartment(ctx kratosx.Context, in *biz.AllDepartmentRequest) ([]*biz.Department, error) {
	var list []*biz.Department

	db := ctx.DB().Model(biz.Department{})
	if in != nil && in.Ids != nil {
		db = db.Where("id in ?", in.Ids)
	}
	return list, db.Find(&list).Error
}

func (r repo) GetParentIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(biz.DepartmentClosure{}).
		Select("parent").
		Where("children=?", rid).
		Scan(&ids).Error
}

func (r repo) GetChildrenIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	var ids = make([]uint32, 0)
	if err := ctx.DB().Model(biz.DepartmentClosure{}).
		Select("children").
		Where("parent=?", rid).
		Scan(&ids).Error; err != nil {
		return ids, err
	}
	ids = append(ids, rid)
	return ids, nil
}

func (r repo) AddDepartment(ctx kratosx.Context, in *biz.Department) (uint32, error) {
	return in.ID(), ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in).Error; err != nil {
			return err
		}

		dcs := r.departmentClosure(ctx, in.Parent(), in.ID())
		if len(dcs) != 0 {
			return tx.Create(&dcs).Error
		}
		return nil
	})
}

func (r repo) departmentClosure(ctx kratosx.Context, pid uint32, id uint32) []biz.DepartmentClosure {
	rcs := []biz.DepartmentClosure{{
		Parent:   pid,
		Children: id,
	}}
	list, _ := r.GetParentIds(ctx, pid)
	for _, item := range list {
		rcs = append(rcs, biz.DepartmentClosure{
			Parent:   item,
			Children: id,
		})
	}
	return rcs
}

func (r repo) UpdateDepartment(ctx kratosx.Context, in *biz.Department) error {
	old, err := r.GetDepartment(ctx, in.ID())
	if err != nil {
		return err
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		// 存在层级变更
		if old.ParentId != in.ParentId {
			if err := tx.Delete(biz.DepartmentClosure{}, "children=?", in.ID()).Error; err != nil {
				return err
			}

			dcs := r.departmentClosure(ctx, in.Parent(), in.ID())
			if len(dcs) != 0 {
				return tx.Create(&dcs).Error
			}
		}
		return tx.Updates(in).Error
	})
}

func (r repo) DeleteDepartment(ctx kratosx.Context, id uint32) error {
	ids, err := r.GetChildrenIds(ctx, id)
	if err != nil {
		return err
	}
	return ctx.DB().Where("id in ?", ids).Delete(biz.Department{}).Error
}

// ParentStatus 获取指定角色的父角色状态
func (r repo) ParentStatus(ctx kratosx.Context, id uint32) bool {
	ids, _ := r.GetParentIds(ctx, id)
	total := int64(0)
	ctx.DB().Where("id in ? and status=false", ids).Count(&total)
	return total == 0
}

func (r repo) GetScope(ctx kratosx.Context, uid, oid uint32) ([]string, error) {
	depIds, err := r.AllManagerDepartmentIds(ctx, uid)
	if err != nil {
		return nil, err
	}

	var value []string
	return value, ctx.DB().Model(biz.DepartmentObject{}).
		Select("value").
		Where("object_id=? and department_id in ?", oid, depIds).
		Scan(&value).Error
}

func (r repo) AllDepartmentObjectValue(ctx kratosx.Context, did, oid uint32) ([]string, error) {
	var value []string
	return value, ctx.DB().Model(biz.DepartmentObject{}).
		Select("value").
		Where("object_id=? and department_id=?", oid, did).
		Scan(&value).Error
}

func (r repo) ImportDepartmentObject(ctx kratosx.Context, list []*biz.DepartmentObject) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		depId := list[0].DepartmentId
		objId := list[0].ObjectId
		if err := tx.Where("department_id=? and object_id=?", depId, objId).
			Delete(biz.DepartmentObject{}).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})
}

func (r repo) AddDepartmentObject(ctx kratosx.Context, in *biz.DepartmentObject) (uint32, error) {
	return in.ID, ctx.DB().Model(biz.DepartmentObject{}).Create(in).Error
}

func (r repo) DeleteDepartmentObjectValue(ctx kratosx.Context, oid uint32, value string) error {
	return ctx.DB().Where("object_id=? and value=?", oid, value).Delete(biz.DepartmentObject{}).Error
}
