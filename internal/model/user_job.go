package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

type UserJob struct {
	types.CreateModel
	JobID  uint32 `json:"job_id" gorm:"not null;comment:角色id"`
	UserID uint32 `json:"user_id" gorm:"not null;comment:菜单id"`
	Job    *Job   `json:"job" gorm:"constraint:OnDelete:cascade"`
}

func (ur *UserJob) FindByUserAndJob(ctx kratosx.Context, userId, jobId uint32) error {
	return ctx.DB().First(ur, "user_id=? and job_id=?", userId, jobId).Error
}

func (ur *UserJob) Add(ctx kratosx.Context) error {
	return ctx.DB().Create(&ur).Error
}

// Update 批量更新角色所属菜单
func (ur *UserJob) Update(ctx kratosx.Context, userId uint32, jobIds []uint32) error {
	// 组装新的菜单数据
	list := make([]UserJob, 0)
	for _, menuId := range jobIds {
		list = append(list, UserJob{
			UserID: userId,
			JobID:  menuId,
		})
	}

	// 删除之后再重新创建
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id=?", userId).Delete(ur).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})
}

// UserJobs 通过角色ID获取角色菜单
func (ur *UserJob) UserJobs(ctx kratosx.Context, userId uint32) ([]*UserJob, error) {
	var list []*UserJob
	return list, ctx.DB().Preload("Job").Find(&list, "user_id=?", userId).Error
}

// DeleteByID 通过id删除
func (ur *UserJob) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(ur, id).Error
}
