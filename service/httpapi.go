package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/giansalex/pluralsight-transcript/model"
)

const transcriptPath = "%s/library/coursetranscripts/%s/%s"

type HttpApiService struct {
	path  string
	token string
}

func NewApi(path string, token string) *HttpApiService {
	return &HttpApiService{
		path:  path,
		token: token,
	}
}

func (api *HttpApiService) GetByCourse(courseID string, lang string) (*model.CourseTranscript, error) {
	url := fmt.Sprintf(transcriptPath, api.path, courseID, lang)
	json, err := api.getContent(url)
	if err != nil {
		return nil, err
	}

	course, err := model.UnmarshalCourse(json)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (api *HttpApiService) getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("ps-jwt", api.token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
