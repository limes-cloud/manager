package menu

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/manager/internal/biz/menu"
	"github.com/limes-cloud/manager/internal/biz/role"
)

func NewRepo() biz.Repo {
	return &repo{}
}

type repo struct {
}

func (r repo) InitBasicMenu(ctx kratosx.Context) {
	var list []*biz.Menu
	if err := ctx.DB().Model(biz.Menu{}).Find(&list, "type=?", biz.MenuBasic).Error; err != nil {
		ctx.Logger().Errorf("init basic api error %s", err.Error())
		return
	}
	for _, item := range list {
		if item.Method != nil && item.Api != nil {
			ctx.Authentication().AddWhitelist(*item.Api, *item.Method)
		}
	}
}

func (r repo) GetMenuByApi(ctx kratosx.Context, api, method string) (*biz.Menu, error) {
	var menu biz.Menu
	return &menu, ctx.DB().First(&menu, "api=? and method=?", api, method).Error
}

func (r repo) AllRoleKeyword(ctx kratosx.Context, id uint32) ([]string, error) {
	var keys []string
	return keys, ctx.DB().Model(role.Role{}).Select("keyword").
		Where("id in (select role_id from role_menu where menu_id = ?)", id).
		Scan(&keys).
		Error
}

func (r repo) AllMenu(ctx kratosx.Context, in *biz.AllMenuRequest) ([]*biz.Menu, error) {
	var list []*biz.Menu

	db := ctx.DB().Model(biz.Menu{})
	if in != nil {
		if in.NotType != nil {
			db = db.Where("type!=?", in.NotType)
		}
		if in.RoleId != nil {
			db = db.Where("id in (select menu_id from role_menu where role_id = ?)", *in.RoleId)
		}
	}

	return list, db.Find(&list).Error
}

func (r repo) GetMenu(ctx kratosx.Context, id uint32) (*biz.Menu, error) {
	var menu biz.Menu
	return &menu, ctx.DB().First(&menu, id).Error
}

func (r repo) AddMenu(ctx kratosx.Context, in *biz.Menu) (uint32, error) {
	return in.ID(), ctx.DB().Create(in).Error
}

func (r repo) UseHome(ctx kratosx.Context, app string, id uint32) error {
	return ctx.DB().Model(biz.Menu{}).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("app=?", app).Where("id!=?", id).Update("is_home", false).Error; err != nil {
			return err
		}
		return tx.Where("app=?", app).Where("id=?", id).Update("is_home", true).Error
	})
}

func (r repo) UpdateMenu(ctx kratosx.Context, in *biz.Menu) error {
	return ctx.DB().Updates(in).Error
}

func (r repo) DeleteMenu(ctx kratosx.Context, ids []uint32) error {
	return ctx.DB().Where("id in ?", ids).Delete(biz.Menu{}).Error
}
