package logic

import (
	v1 "manager/api/v1"
	"manager/config"
	"manager/consts"
	"manager/internal/model"
	"manager/pkg/md"
	"manager/pkg/util"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/kratos"
)

type Department struct {
	conf *config.Config
}

func NewDepartment(conf *config.Config) *Department {
	return &Department{
		conf: conf,
	}
}

func (d *Department) Get(ctx kratos.Context, in *v1.GetDepartmentRequest) (*v1.GetDepartmentReply, error) {
	department := model.Department{}
	if err := department.OneByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	reply := v1.GetDepartmentReply{}
	if err := util.Transform(department, &reply.Department); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	return &reply, nil
}

// Tree 查询部门树
func (d *Department) Tree(ctx kratos.Context) (*v1.GetDepartmentTreeReply, error) {
	department := model.Department{}
	tree, err := department.Tree(ctx)
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	reply := v1.GetDepartmentTreeReply{}
	if err = util.Transform(tree, &reply.Department); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}
	return &reply, nil
}

func (d *Department) Add(ctx kratos.Context, in *v1.AddDepartmentRequest) (*emptypb.Empty, error) {
	// 判断用户是否具有部门的权限
	user := model.User{}
	depIds, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}
	if !util.InList(depIds, in.ParentId) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	department := model.Department{}
	if err := util.Transform(in, &department); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	if err := department.Create(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

func (d *Department) Update(ctx kratos.Context, in *v1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorEditSystemData()
	}

	if in.Id == md.DepartmentId(ctx) {
		return nil, v1.ErrorDeleteSelfDepartment()
	}

	oldDep := model.Department{}
	if err := oldDep.OneByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorNotFound()
	}

	// 判断用户是否具有部门的权限
	user := model.User{}
	depIds, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 判断是否有移动部门的权限
	if oldDep.ParentID != in.ParentId && !util.InList(depIds, in.ParentId) {
		return nil, v1.ErrorDepartmentPermissions()
	}
	// 判断是否有当前部门的权限
	if !util.InList(depIds, in.Id) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	department := model.Department{}
	if err := util.Transform(in, &department); err != nil {
		return nil, v1.ErrorTransformFormat(err.Error())
	}

	if err := department.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}

func (d *Department) Delete(ctx kratos.Context, in *v1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	if in.Id == consts.SuperAdmin {
		return nil, v1.ErrorDeleteSystemData()
	}
	// 判断用户是否具有部门的权限
	user := model.User{}
	depIds, err := user.DataScope(ctx, md.UserId(ctx))
	if err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	if !util.InList(depIds, in.Id) {
		return nil, v1.ErrorDepartmentPermissions()
	}

	department := model.Department{}
	if err := department.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	return nil, nil
}
