package department

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"github.com/limes-cloud/kratosx/pkg/util"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/biz/object"
	"github.com/limes-cloud/manager/internal/config"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	repo       Repo
	objectRepo object.Repo
	conf       *config.Config
}

func NewUseCase(conf *config.Config, repo Repo, objectRepo object.Repo) *UseCase {
	return &UseCase{
		repo:       repo,
		objectRepo: objectRepo,
		conf:       conf,
	}
}

// DepartmentTree 查询当前用户的部门树
func (u *UseCase) DepartmentTree(ctx kratosx.Context) ([]tree.Tree, error) {
	// 查询子部门id列表信息
	list, err := u.repo.AllManagerDepartment(ctx, md.UserId(ctx))
	if err != nil {
		return nil, err
	}

	// 构建树枝数组
	var ts []tree.Tree
	for _, item := range list {
		ts = append(ts, item)
	}

	// 生成树
	return tree.BuildArrayTree(ts), nil
}

func (u *UseCase) AddDepartment(ctx kratosx.Context, in *Department) (uint32, error) {
	ids, err := u.repo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return 0, errors.Database()
	}
	if !util.InList(ids, in.ParentId) {
		return 0, errors.DepartmentPermissions()
	}

	// 创建部门信息
	id, err := u.repo.AddDepartment(ctx, in)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}

	return id, nil
}

func (u *UseCase) UpdateDepartment(ctx kratosx.Context, in *Department) error {
	// 超级管理员不允许修改
	if in.ID() == 1 {
		return errors.EditSystemData()
	}

	//  不能修改当前部门
	departmentId := md.DepartmentId(ctx)
	if in.ID() == md.DepartmentId(ctx) {
		return errors.DepartmentPermissions()
	}

	// 判断当前部门是否具有权限
	ids, err := u.repo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	if !util.InList(ids, in.ID()) {
		return errors.DepartmentPermissions()
	}

	// 查询历史部门信息
	oldDepartment, err := u.repo.GetDepartment(ctx, departmentId)
	if err != nil {
		return errors.NotFound()
	}

	// 如果存在移动部门，判断是否有父级部门的权限
	if oldDepartment.ParentId != in.ParentId && !util.InList(ids, in.ParentId) {
		return errors.DepartmentPermissions()
	}

	// 更新部门信息
	if err := u.repo.UpdateDepartment(ctx, in); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	return nil
}

func (u *UseCase) DeleteDepartment(ctx kratosx.Context, id uint32) error {
	// 超级管理员不允许删除
	if id == 1 {
		return errors.DeleteSystemData()
	}

	// 不能删除当前部门
	departmentId := md.DepartmentId(ctx)
	if id == departmentId {
		return errors.DeleteSelfDepartment()
	}

	// 是否具有部门管理权限
	ids, err := u.repo.AllManagerDepartmentIds(ctx, md.UserId(ctx))
	if err != nil {
		return errors.Database()
	}
	if !util.InList(ids, id) {
		return errors.DepartmentPermissions()
	}

	// 删除部门
	if err := u.repo.DeleteDepartment(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	return nil
}

// AllDepartmentObjectValue 获取指定部门的指定资源对象的全部值
func (u *UseCase) AllDepartmentObjectValue(ctx kratosx.Context, in *AllDepartmentObjectValueRequest) ([]string, error) {
	list, err := u.repo.AllDepartmentObjectValue(ctx, in.DepartmentId, in.ObjectId)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return list, nil
}

// ImportDepartmentObject 导入部门资源对象
func (u *UseCase) ImportDepartmentObject(ctx kratosx.Context, in []*DepartmentObject) error {
	if err := u.repo.ImportDepartmentObject(ctx, in); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// AddDepartmentObject 添加部门资源对象
func (u *UseCase) AddDepartmentObject(ctx kratosx.Context, objectKey, value string) (uint32, error) {
	obj, err := u.objectRepo.GetObjectByKeyword(ctx, objectKey)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}

	id, err := u.repo.AddDepartmentObject(ctx, &DepartmentObject{
		ObjectId:     obj.ID,
		DepartmentId: md.DepartmentId(ctx),
		Value:        value,
	})
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}

	return id, nil
}

// DeleteDepartmentObject 删除部门资源对象
func (u *UseCase) DeleteDepartmentObject(ctx kratosx.Context, objectKey, value string) error {
	obj, err := u.objectRepo.GetObjectByKeyword(ctx, objectKey)
	if err != nil {
		return errors.DatabaseFormat(err.Error())
	}

	if err := u.repo.DeleteDepartmentObjectValue(ctx, obj.ID, value); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
