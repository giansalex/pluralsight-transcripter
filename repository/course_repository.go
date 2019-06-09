package repository

import (
	"github.com/giansalex/pluralsight-transcript/model"
	"github.com/jinzhu/gorm"
)

type DbCourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *DbCourseRepository {
	return &DbCourseRepository{
		db: db,
	}
}

func (repository *DbCourseRepository) List() []model.Course {
	var courses []model.Course
	repository.db.Find(&courses)

	return courses
}
