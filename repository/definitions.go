package repository

import "github.com/giansalex/pluralsight-transcript/model"

type UserRepository interface {
	GetUser() *model.User
}

type TranscriptRepository interface {
	Save(course *model.CourseTranscript)
}

type CourseRepository interface {
	List() []model.Course
}
