package test

import (
	"io/ioutil"
	"testing"

	"github.com/giansalex/pluralsight-transcript/model"
)

func TestDeserializeJson(t *testing.T) {
	buf, _ := ioutil.ReadFile("test.json")

	course, _ := model.UnmarshalCourse(buf)

	if course == nil {
		t.Error("Null Course")
	}

	if len(course.Modules) == 0 {
		t.Error("Empty Modules")
	}
}
