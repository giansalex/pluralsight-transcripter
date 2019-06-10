package service

import (
	"fmt"

	"github.com/giansalex/pluralsight-transcript/repository"
)

type DefaultTranscriptService struct {
	courses     repository.CourseRepository
	transcripts repository.TranscriptRepository
	api         ApiService
}

func NewTranscriptService(courses repository.CourseRepository, transcripts repository.TranscriptRepository, api ApiService) *DefaultTranscriptService {
	return &DefaultTranscriptService{
		courses:     courses,
		transcripts: transcripts,
		api:         api,
	}
}

func (service *DefaultTranscriptService) Transcript(lang string) error {
	courses := service.courses.List()

	service.transcripts.ClearAll()
	for _, course := range courses {
		fmt.Println("Transcript:", course.Title.String)
		courseTranscript, err := service.api.GetByCourse(course.Name.String, lang)

		if err != nil {
			return err
		}

		service.transcripts.Save(courseTranscript)
	}

	return nil
}
