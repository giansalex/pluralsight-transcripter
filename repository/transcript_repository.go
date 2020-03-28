package repository

import (
	"database/sql"

	"github.com/giansalex/pluralsight-transcripter/model"
	"github.com/jinzhu/gorm"
)

type DbTranscriptRepository struct {
	db *gorm.DB
}

func NewTranscriptRepository(db *gorm.DB) *DbTranscriptRepository {
	return &DbTranscriptRepository{
		db: db,
	}
}

func (repository *DbTranscriptRepository) ClearAll() {
	repository.db.Delete(model.Cliptranscript{})
}

func (repository *DbTranscriptRepository) Save(course *model.CourseTranscript) {

	for _, module := range course.Modules {
		for _, clipTranscript := range module.ClipTranscripts {
			var clip model.Clip
			clipName := sql.NullString{String: clipTranscript.ClipID, Valid: true}
			repository.db.Where(&model.Clip{Name: clipName}).First(&clip)

			for _, transcript := range clipTranscript.Transcripts {
				newTranscript := model.Cliptranscript{
					ClipID:    clip.ID,
					StartTime: sql.NullInt64{Int64: transcript.RelativeStartTime, Valid: true},
					EndTime:   sql.NullInt64{Int64: transcript.RelativeEndTime, Valid: true},
					Text:      sql.NullString{String: transcript.Text, Valid: true},
				}

				repository.db.Create(&newTranscript)
			}
		}
	}
}
