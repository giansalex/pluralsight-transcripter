package service

import "github.com/giansalex/pluralsight-transcript/model"

type ApiService interface {
	GetByCourse(courseID string, lang string) (*model.CourseTranscript, error)
}

type TranscriptService interface {
	Transcript(lang string) error
}
