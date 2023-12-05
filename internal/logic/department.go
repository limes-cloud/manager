package logic

import (
	v1 "manager/api/v1"
	"manager/config"
	"manager/consts"
	"manager/internal/model"
	"manager/pkg/md"
	"manager/pkg/util"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Department struct {
	conf *config.Config
}

func NewDepartment(conf *config.Config) *Department {
	return &Department{
		conf: conf,
	}
}

func (d *Department) Get(ctx kratosx.Context, in *v1.GetDepartmentRequest) (*v1.GetDepartmentReply, error) {
	department := model.Department{}
	if err := department.OneByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.GetDepartmentReply{}
	if err := util.Transform(department, &reply.Department); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Tree 查询部门树
func (d *Department) Tree(ctx kratosx.Context) (*v1.GetDepartmentTreeReply, error) {
	department := model.Department{}
	tree, err := department.Tree(ctx)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.GetDepartmentTreeReply{}
	if err = util.Transform(tree, &reply.Department); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (d *Department) Add(ctx kratosx.Context, in *v1.AddDepartmentRequest) (*emptypb.Empty, error) {
	// 判断用户是否具有部门的权限
	user := model.User{}
	depIds, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	if !util.InList(depIds, in.ParentId) {
		return nil, v1.DepartmentPermissionsError()
	}

	department := model.Department{}
	if err := util.Transform(in, &department); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := department.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

func (d *Department) Update(ctx kratosx.Context, in *v1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.EditSystemDataError()
	}

	if in.Id == md.DepartmentId(ctx) {
		return nil, v1.DeleteSelfDepartmentError()
	}

	oldDep := model.Department{}
	if err := oldDep.OneByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundError()
	}

	// 判断用户是否具有部门的权限
	user := model.User{}
	depIds, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 判断是否有移动部门的权限
	if oldDep.ParentID != in.ParentId && !util.InList(depIds, in.ParentId) {
		return nil, v1.DepartmentPermissionsError()
	}
	// 判断是否有当前部门的权限
	if !util.InList(depIds, in.Id) {
		return nil, v1.DepartmentPermissionsError()
	}

	department := model.Department{}
	if err := util.Transform(in, &department); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := department.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

func (d *Department) Delete(ctx kratosx.Context, in *v1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.DeleteSystemDataError()
	}
	// 判断用户是否具有部门的权限
	user := model.User{}
	depIds, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	if !util.InList(depIds, in.Id) {
		return nil, v1.DepartmentPermissionsError()
	}

	department := model.Department{}
	if err := department.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}
