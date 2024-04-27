package user

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/manager/internal/biz/user"
)

func NewRepo() biz.Repo {
	return &repo{}
}

type repo struct {
}

func (r repo) GetUserIdsByDepartmentIds(ctx kratosx.Context, ids []uint32) ([]uint32, error) {
	var list []uint32
	return list, ctx.DB().Model(biz.User{}).
		Select("id").
		Where("department_id in ?", ids).
		Scan(&list).Error
}

func (r repo) HasRole(ctx kratosx.Context, uid, rid uint32) bool {
	return ctx.DB().Where("user_id=? and role_id=?", uid, rid).First(&biz.UserRole{}).RowsAffected != 0
}

func (r repo) GetUser(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Role").
		Preload("Department").
		Preload("Roles").
		Preload("Jobs").
		First(&user, id).Error
}

func (r repo) GetUserByPhone(ctx kratosx.Context, phone string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Role").
		Preload("Department").
		Preload("Roles").
		Preload("Jobs").
		First(&user, "phone=?", phone).Error
}

func (r repo) GetUserByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Role").
		Preload("Department").
		Preload("Roles").
		Preload("Jobs").
		First(&user, "email=?", email).Error
}

func (r repo) PageUser(ctx kratosx.Context, in *biz.PageUserRequest) ([]*biz.User, uint32, error) {
	var list []*biz.User
	var total int64

	db := ctx.DB().Model(biz.User{}).
		Preload("Role").
		Preload("Department").
		Preload("Roles").
		Preload("Jobs")

	if in.Username != nil {
		db = db.Where("email=? or phone=?", *in.Username, *in.Username)
	}
	if in.Status != nil {
		db = db.Where("status=?", *in.Status)
	}
	if in.RoleId != nil {
		db = db.Where("role_id=?", *in.RoleId)
	}
	if in.DepartmentId != nil {
		db = db.Where("department_id=?", *in.DepartmentId)
	}
	if in.Name != nil {
		db = db.Where("name like ?", "%"+*in.Name+"%")
	}
	if in.StartTime != nil {
		db = db.Where("created_at >= ?", *in.StartTime)
	}
	if in.EndTime != nil {
		db = db.Where("created_at <= ?", *in.EndTime)
	}
	if in.DepartmentIds != nil {
		db = db.Where("department_id in ?", in.DepartmentIds)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (r repo) AddUser(ctx kratosx.Context, in *biz.User) (uint32, error) {
	return in.ID, ctx.DB().Create(in).Error
}

func (r repo) UpdateUser(ctx kratosx.Context, in *biz.User) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if len(in.UserRoles) != 0 {
			if err := tx.Where("user_id=?", in.ID).Delete(biz.UserRole{}).Error; err != nil {
				return err
			}
		}

		if len(in.UserJobs) != 0 {
			if err := tx.Where("user_id=?", in.ID).Delete(biz.UserJob{}).Error; err != nil {
				return err
			}
		}

		return tx.Updates(in).Error
	})
}

func (r repo) DeleteUser(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.User{}, id).Error
}
