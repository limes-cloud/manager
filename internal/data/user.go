package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	file "github.com/limes-cloud/resource/api/resource/file/v1"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	biz "github.com/limes-cloud/manager/internal/biz/user"
	"github.com/limes-cloud/manager/internal/data/model"
	"github.com/limes-cloud/manager/internal/pkg/service"
)

type userRepo struct {
	departmentRepo // fixed code
	roleRepo
}

func NewUserRepo() biz.Repo {
	return &userRepo{}
}

// ToUserEntity model转entity
func (r userRepo) ToUserEntity(ctx kratosx.Context, m *model.User) *biz.User {
	e := &biz.User{}
	_ = valx.Transform(m, e)

	if e.Avatar != nil && *e.Avatar != "" {
		e.Avatar = proto.String(r.GetResourceURL(ctx, *e.Avatar))
	}
	return e
}

// ToUserModel entity转model
func (r userRepo) ToUserModel(e *biz.User) *model.User {
	m := &model.User{}
	_ = valx.Transform(e, m)
	return m
}

// GetUserByPhone 获取指定数据 fixed code
func (r userRepo) GetUserByPhone(ctx kratosx.Context, phone string) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"id", "department_id", "role_id", "name", "nickname", "gender", "avatar", "phone", "email", "status", "setting", "logged_at", "created_at", "updated_at"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Roles").Preload("Jobs").Preload("Department")
	if err := db.Where("phone = ?", phone).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(ctx, &m), nil
}

// GetUserByEmail 获取指定数据 fixed code
func (r userRepo) GetUserByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"id", "department_id", "role_id", "name", "nickname", "gender", "avatar", "phone", "email", "status", "setting", "logged_at", "created_at", "updated_at"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Roles").Preload("Jobs").Preload("Department")
	if err := db.Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(ctx, &m), nil
}

// GetUser 获取指定的数据 fixed code
func (r userRepo) GetUser(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"id", "department_id", "role_id", "name", "nickname", "gender", "avatar", "phone", "email", "status", "setting", "logged_at", "created_at", "updated_at"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Roles").Preload("Jobs").Preload("Department")
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(ctx, &m), nil
}

// GetUserToken 获取指定的用户token
func (r userRepo) GetUserToken(ctx kratosx.Context, id uint32) (*string, int64, error) {
	var (
		m  = model.User{}
		fs = []string{"token", "logged_at"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, 0, err
	}
	return m.Token, m.LoggedAt, nil
}

// GetUserPassword 获取用户的密码
func (r userRepo) GetUserPassword(ctx kratosx.Context, id uint32) (string, error) {
	var password string
	db := ctx.DB().Select("password")
	if err := db.Model(model.User{}).Where("id=?", id).Scan(&password).Error; err != nil {
		return "", err
	}
	return password, nil
}

// ListUser 获取列表 fixed code
func (r userRepo) ListUser(ctx kratosx.Context, req *biz.ListUserRequest) ([]*biz.User, uint32, error) {
	var (
		bs    []*biz.User
		ms    []*model.User
		total int64
		fs    = []string{"id", "department_id", "role_id", "name", "nickname", "gender", "avatar", "phone", "email", "status", "setting", "logged_at", "created_at", "updated_at"}
	)

	db := ctx.DB().Model(model.User{}).Select(fs).Preload("Role").Preload("Department")

	if req.DepartmentId != nil {
		db = db.Where("department_id = ?", *req.DepartmentId)
	}
	if req.RoleId != nil {
		db = db.Where("role_id = ?", *req.RoleId)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Phone != nil {
		db = db.Where("phone = ?", *req.Phone)
	}
	if req.Email != nil {
		db = db.Where("email = ?", *req.Email)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if len(req.LoggedAts) == 2 {
		db = db.Where("logged_at BETWEEN ? AND ?", req.LoggedAts[0], req.LoggedAts[1])
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}
	if req.DepartmentIds != nil {
		db = db.Where("department_id in ?", req.DepartmentIds)
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
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	if err := db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: *req.OrderBy},
		Desc:   *req.Order == "desc",
	}).Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToUserEntity(ctx, m))
	}
	return bs, uint32(total), nil
}

// CreateUser 创建数据
func (r userRepo) CreateUser(ctx kratosx.Context, req *biz.User) (uint32, error) {
	m := r.ToUserModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateUser 更新数据
func (r userRepo) UpdateUser(ctx kratosx.Context, req *biz.User) error {
	return ctx.DB().Updates(r.ToUserModel(req)).Error
}

// UpdateUserStatus 更新数据状态
func (r userRepo) UpdateUserStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.User{}).Where("id=?", id).Update("status", status).Error
}

// DeleteUser 删除数据
func (r userRepo) DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.User{})
	return uint32(db.RowsAffected), db.Error
}

// HasUserDataScope 是否具有用户权限
func (r userRepo) HasUserDataScope(ctx kratosx.Context, pid, uid uint32) (bool, error) {
	all, scopes, err := r.GetDepartmentDataScope(ctx, pid)
	if err != nil {
		return false, err
	}
	if all {
		return true, nil
	}

	var depId uint32
	if err = ctx.DB().Model(model.User{}).
		Select("department_id").
		Where("id=?", uid).
		Scan(&depId).Error; err != nil {
		return false, err
	}

	return valx.InList(scopes, depId), nil
}

func (r userRepo) GetResourceURL(ctx kratosx.Context, sha string) string {
	client, err := service.NewResourceFile(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "init resource client error", "err", err.Error())
		return ""
	}
	reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: proto.String(sha)})
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource sha error", "err", err.Error())
		return ""
	}

	return reply.Url
}
