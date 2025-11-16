package code

import (
	"golang-api/database"
	"golang-api/query"
	"time"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type ICodeModel interface {
	common.IModel[*Code]
	QueryFindAll(q query.QueryFilter) ([]*Code, error)
	FindOneByCodeAndEmail(code, email string) (*Code, error)
	DeleteExpiredCodes() error
}

type CodeModel struct {
	*common.Model[*Code]
	databaseService database.IDatabaseService
}

func NewCodeModel(module core.IModule) *CodeModel {
	service := &CodeModel{
		Model:           common.NewModel[*Code]("CodeModel"),
		databaseService: module.Get("DatabaseService").(database.IDatabaseService),
	}

	module.On("db:migrate", service.Migrate)

	return service
}

func (cm *CodeModel) OnInit() error {
	cm.SetDB(cm.databaseService.GetDB())
	return nil
}

func (cm *CodeModel) QueryFindAll(q query.QueryFilter) ([]*Code, error) {
	var items []*Code

	tx := cm.databaseService.GetDB().Model(&Code{}).
		Offset(q.GetSkip()).
		Where(q.GetWhere()).
		Order(q.GetSort())

	if q.GetLimit() != 0 {
		tx.Limit(q.GetLimit())
	}

	err := tx.Find(&items).Error
	return items, err
}

func (cm *CodeModel) FindOneByCodeAndEmail(code, email string) (*Code, error) {
	var item *Code
	err := cm.databaseService.GetDB().Model(&Code{}).
		Where("code = ? AND email = ?", code, email).
		First(&item).Error
	return item, err
}

func (cm *CodeModel) DeleteExpiredCodes() error {
	return cm.databaseService.GetDB().Model(&Code{}).
		Where("expires_at < ?", time.Now()).
		Delete(&Code{}).Error
}
