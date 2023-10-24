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



func (mdl *model) Paginate(page, size int) []goly.Goly {
	var goly []goly.Goly

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&goly)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return goly
}

func (mdl *model) Insert(newGoly *goly.Goly) *goly.Goly {
	result := mdl.db.Create(newGoly)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	return newGoly
}

func (mdl *model) SelectByID(golyID int) *goly.Goly {
	var goly goly.Goly
	result := mdl.db.First(&goly, golyID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &goly
}

func (mdl *model) SearchingGoly(short string) ([]goly.Goly, error) {
    var golies []goly.Goly
    result := mdl.db.Where("goly LIKE ?", "%"+short+"%").Find(&golies)
    if result.Error != nil {
        log.Error(result.Error)
        return nil, result.Error
    }
    return golies, nil
}

func(mdl *model)FindByGolyUrl(url string) (goly.Goly, error) {
	var goly goly.Goly
    result := mdl.db.Where("goly = ?", url).First(&goly)

    if result.Error!= nil {
        log.Error(result.Error)
        return goly, result.Error
    }

    return goly, nil
}

func (mdl *model) Update(goly goly.Goly) int64 {
	result := mdl.db.Updates(&goly)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) UpdateButton(goly goly.Goly) error {
    result := mdl.db.Updates(&goly)

    if result.Error != nil {
        log.Error(result.Error)
        return result.Error
    }

    return nil
}

func (mdl *model) DeleteByID(golyID int) int64 {
	result := mdl.db.Delete(&goly.Goly{}, golyID)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}