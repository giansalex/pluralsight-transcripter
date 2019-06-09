package model

import "encoding/json"

func UnmarshalCourse(data []byte) (*CourseTranscript, error) {
	var r CourseTranscript
	err := json.Unmarshal(data, &r)
	return &r, err
}

type CourseTranscript struct {
	CourseID string             `json:"courseId"`
	Modules  []ModuleTranscript `json:"modules"`
}

type ModuleTranscript struct {
	ModuleID        string              `json:"moduleId"`
	ModuleTitle     string              `json:"moduleTitle"`
	ClipTranscripts []ClipTranscription `json:"clipTranscripts"`
}

type ClipTranscription struct {
	ClipID       string       `json:"clipId"`
	LanguageCode string       `json:"languageCode"`
	Transcripts  []Transcript `json:"transcripts"`
}

type Transcript struct {
	RelativeStartTime int64  `json:"relativeStartTime"`
	RelativeEndTime   int64  `json:"relativeEndTime"`
	Text              string `json:"text"`
}
