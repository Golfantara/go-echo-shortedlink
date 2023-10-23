package repository

import (
	"shortlink/features/goly"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) goly.Repository {
	return &model{
        db: db,
    }
}

func (mdl *model) Insert(newGoly *goly.Goly) *goly.Goly {
	result := mdl.db.Create(newGoly)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	return newGoly
}

func (mdl *model) SelectByID(userID int) *goly.Goly {
	var goly goly.Goly
	result := mdl.db.First(&goly, userID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &goly
}

func (mdl *model) Update(goly goly.Goly) int64 {
	result := mdl.db.Updates(&goly)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(golyID int) int64 {
	result := mdl.db.Delete(&goly.Goly{}, golyID)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}