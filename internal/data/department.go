package data

import (
	"errors"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	biz "github.com/limes-cloud/manager/internal/biz/department"
	"github.com/limes-cloud/manager/internal/data/model"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
	"strings"
)

type departmentRepo struct {
}

func NewDepartmentRepo() biz.Repo {
	return &departmentRepo{}
}

// ToDepartmentEntity model转entity
func (r departmentRepo) ToDepartmentEntity(m *model.Department) *biz.Department {
	e := &biz.Department{}
	_ = valx.Transform(m, e)
	return e
}

// ToDepartmentModel entity转model
func (r departmentRepo) ToDepartmentModel(e *biz.Department) *model.Department {
	m := &model.Department{}
	_ = valx.Transform(e, m)
	return m
}

// ListDepartment 获取列表 fixed code
func (r departmentRepo) ListDepartment(ctx kratosx.Context, req *biz.ListDepartmentRequest) ([]*biz.Department, uint32, error) {
	var (
		bs    []*biz.Department
		ms    []*model.Department
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Department{}).Select(fs)

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	if err := db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: *req.OrderBy},
		Desc:   *req.Order == "desc",
	}).Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToDepartmentEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateDepartment 创建数据
func (r departmentRepo) CreateDepartment(ctx kratosx.Context, req *biz.Department) (uint32, error) {
	m := r.ToDepartmentModel(req)
	return m.Id, ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(m).Error; err != nil {
			return err
		}
		return r.appendDepartmentChildren(ctx, req.ParentId, m.Id)
	})
}

// GetDepartment 获取指定的数据
func (r departmentRepo) GetDepartment(ctx kratosx.Context, id uint32) (*biz.Department, error) {
	var (
		m  = model.Department{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToDepartmentEntity(&m), nil
}

// UpdateDepartment 更新数据
func (r departmentRepo) UpdateDepartment(ctx kratosx.Context, req *biz.Department) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := r.GetDepartment(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.ParentId != req.ParentId {
			if err := r.removeDepartmentParent(ctx, req.Id); err != nil {
				return err
			}
			if err := r.appendDepartmentChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(r.ToDepartmentModel(req)).Error
	})
}

// DeleteDepartment 删除数据
func (r departmentRepo) DeleteDepartment(ctx kratosx.Context, ids []uint32) (uint32, error) {
	var del []uint32
	for _, id := range ids {
		del = append(del, id)
		childrenIds, err := r.GetDepartmentChildrenIds(ctx, id)
		if err != nil {
			return 0, err
		}
		del = append(del, childrenIds...)
	}
	db := ctx.DB().Where("id in ?", del).Delete(&model.Department{})
	return uint32(db.RowsAffected), db.Error
}

// GetDepartmentChildrenIds 获取指定id的所有子id
func (r departmentRepo) GetDepartmentChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.DepartmentClosure{}).
		Select("children").
		Where("parent=?", id).
		Scan(&ids).Error
}

// GetDepartmentParentIds 获取指定id的所有父id
func (r departmentRepo) GetDepartmentParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.DepartmentClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendDepartmentChildren 添加id到指定的父id下
func (r departmentRepo) appendDepartmentChildren(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*model.DepartmentClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := r.GetDepartmentParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &model.DepartmentClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeDepartmentParent 删除指定id的所有父层级
func (r departmentRepo) removeDepartmentParent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&model.DepartmentClosure{}, "children=?", id).Error
}

// GetDepartmentByKeyword 获取指定数据
func (r departmentRepo) GetDepartmentByKeyword(ctx kratosx.Context, keyword string) (*biz.Department, error) {
	var (
		m  = model.Department{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToDepartmentEntity(&m), nil
}

// GetDepartmentDataScope 获取部门权限
func (r departmentRepo) GetDepartmentDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error) {
	user := model.User{}
	if err := ctx.DB().
		Select([]string{"role_id", "department_id"}).
		First(&user, uid).Error; err != nil {
		return false, nil, err
	}

	role := model.Role{}
	if err := ctx.DB().
		Select([]string{"data_scope"}).
		First(&role, "id=?", user.RoleId).Error; err != nil {
		return false, nil, err
	}

	if role.DataScope == biz.DataScopeAll {
		return true, []uint32{}, nil
	}

	if role.DataScope == biz.DataScopeCurrentDown && user.DepartmentId == 1 {
		return true, []uint32{}, nil
	}

	switch role.DataScope {
	case biz.DataScopeAll:
		ids := make([]uint32, 0)
		if err := ctx.DB().Select("id").Model(biz.Department{}).Scan(&ids).Error; err != nil {
			return false, nil, err
		}
		return false, ids, nil
	case biz.DataScopeCurrent:
		return false, []uint32{user.DepartmentId}, nil
	case biz.DataScopeCurrentDown:
		ids, err := r.GetDepartmentChildrenIds(ctx, user.DepartmentId)
		if err != nil {
			return false, nil, err
		}
		ids = append(ids, user.DepartmentId)
		return false, ids, nil
	case biz.DataScopeDown:
		ids, err := r.GetDepartmentChildrenIds(ctx, user.DepartmentId)
		if err != nil {
			return false, nil, err
		}
		return false, ids, nil
	case biz.DataScopeCustom:
		if role.DepartmentIds == nil {
			return false, []uint32{}, nil
		}
		ids := make([]uint32, 0)
		arr := strings.Split(*role.DepartmentIds, ",")
		for _, id := range arr {
			ids = append(ids, valx.ToUint32(id))
		}
		return false, ids, nil
	}
	return false, []uint32{}, nil
}
