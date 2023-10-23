package repository

import (
	"shortlink/features/auth"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.Repository {
	return &model{
        db: db,
    }
}

func (mdl *model) Paginate(page, size int) []auth.Users {
	var user []auth.Users

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&user)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return user
}

func (mdl *model) Insert(newUsers *auth.Users) *auth.Users {
	result := mdl.db.Create(newUsers)

	if result.Error!= nil {
        log.Error(result.Error)
		return nil
    }
	return newUsers
}

func (mdl *model) SelectByID(userID int) *auth.Users {
	var user auth.Users
	result := mdl.db.First(&user, userID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &user
}

func (mdl *model) Update(user auth.Users) int64 {
	result := mdl.db.Updates(&user)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(userID int) int64 {
	result := mdl.db.Delete(&auth.Users{}, userID)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) Login(email string, password string) (*auth.Users, error) {
	var user auth.Users
	result := mdl.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(result.Error)
		return nil, result.Error
	}
	return &user, nil
}