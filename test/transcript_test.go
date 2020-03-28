package test

import (
	"testing"

	"github.com/giansalex/pluralsight-transcripter/model"
	"github.com/giansalex/pluralsight-transcripter/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestSaveCourse(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "pluralsight.db")
	defer db.Close()
	transcriptRepository := repository.NewTranscriptRepository(db)

	transcripts := []model.Transcript{}
	transcripts = append(transcripts, model.Transcript{
		RelativeStartTime: 5070,
		RelativeEndTime:   7458,
		Text:              "Hola a todos",
	})

	transcripts = append(transcripts, model.Transcript{
		RelativeStartTime: 7458,
		RelativeEndTime:   9547,
		Text:              "bienvenidos a mi curso",
	})

	course := model.CourseTranscript{
		CourseID: "edcf503b-3538-427a-96c0-598b0bd79413",
		Modules: []model.ModuleTranscript{
			model.ModuleTranscript{
				ModuleID: "cd7b4491-a75d-4af5-b916-ed044490dee3",
				ClipTranscripts: []model.ClipTranscription{
					model.ClipTranscription{
						ClipID:       "16a0602c-4104-422b-a381-d71d51a8b8ee",
						LanguageCode: "es",
						Transcripts:  transcripts,
					},
				},
			},
		},
	}

	transcriptRepository.Save(&course)
}
