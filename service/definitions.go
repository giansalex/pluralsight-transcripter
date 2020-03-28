package service

import "github.com/giansalex/pluralsight-transcripter/model"

// APIService represent pluralsight API
type APIService interface {
	GetByCourse(courseID string, lang string) (*model.CourseTranscript, error)
}

// TranscriptService service for transcript courses
type TranscriptService interface {
	Transcript(lang string) error
}
