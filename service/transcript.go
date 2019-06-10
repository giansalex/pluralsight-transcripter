package service

import (
	"fmt"

	"github.com/giansalex/pluralsight-transcript/model"
	"github.com/giansalex/pluralsight-transcript/repository"
)

// DefaultTranscriptService implementation TranscriptService
type DefaultTranscriptService struct {
	courses     repository.CourseRepository
	transcripts repository.TranscriptRepository
	api         APIService
}

// NewTranscriptService instance creator
func NewTranscriptService(courses repository.CourseRepository, transcripts repository.TranscriptRepository, api APIService) *DefaultTranscriptService {
	return &DefaultTranscriptService{
		courses:     courses,
		transcripts: transcripts,
		api:         api,
	}
}

// Transcript translate to specify language
func (service *DefaultTranscriptService) Transcript(lang string) error {
	courses := service.courses.List()
	transcripts := []*model.CourseTranscript{}

	for _, course := range courses {
		fmt.Println("Download Transcript:", course.Title.String)
		courseTranscript, err := service.api.GetByCourse(course.Name.String, lang)

		if err != nil {
			return err
		}

		transcripts = append(transcripts, courseTranscript)
	}

	service.transcripts.ClearAll()
	for index, transcript := range transcripts {
		fmt.Println("Save Transcript:", courses[index].Title.String)
		service.transcripts.Save(transcript)
	}

	return nil
}
