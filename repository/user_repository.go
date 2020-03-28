package repository

import (
	"github.com/giansalex/pluralsight-transcripter/model"
	"github.com/jinzhu/gorm"
)

type DbUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *DbUserRepository {
	return &DbUserRepository{
		db: db,
	}
}

func (repository *DbUserRepository) GetUser() *model.User {
	var user model.User
	repository.db.First(&user)

	return &user
}
