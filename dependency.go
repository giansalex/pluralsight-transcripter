package main

import (
	"github.com/giansalex/pluralsight-transcripter/model"
	"github.com/giansalex/pluralsight-transcripter/repository"
	"github.com/giansalex/pluralsight-transcripter/service"
	"go.uber.org/dig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func newUserRepository(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

func newCourseRepository(db *gorm.DB) repository.CourseRepository {
	return repository.NewCourseRepository(db)
}

func newTranscriptRepository(db *gorm.DB) repository.TranscriptRepository {
	return repository.NewTranscriptRepository(db)
}

func getCurrentUser(users repository.UserRepository) *model.User {
	return users.GetUser()
}

func newAPI(user *model.User) service.APIService {
	return service.NewAPI(apiURL, user.Jwt.String)
}

func useTranscriptService(transcript *service.DefaultTranscriptService) service.TranscriptService {
	return transcript
}

func buildContainer(dbPath string) *dig.Container {
	container := dig.New()

	container.Provide(func() (*gorm.DB, error) {
		return gorm.Open("sqlite3", dbPath)
	})
	container.Provide(newUserRepository)
	container.Provide(newCourseRepository)
	container.Provide(newTranscriptRepository)
	container.Provide(getCurrentUser)
	container.Provide(newAPI)
	container.Provide(service.NewTranscriptService)
	container.Provide(useTranscriptService)

	return container
}
