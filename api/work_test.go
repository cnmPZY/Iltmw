package api

import (
	"Iltmw/model"
	"testing"
)

func TestGetAns(t *testing.T) {
	ques := &model.ModelObj{
		PaperId: "OS6hpoUIRVM2BLDxVkX",
		List: []struct {
			PaperDetailId string      `json:"paperDetailId"`
			Title         string      `json:"title"`
			AnswerA       string      `json:"answerA"`
			AnswerB       string      `json:"answerB"`
			AnswerC       string      `json:"answerC"`
			AnswerD       string      `json:"answerD"`
			QuestionId    interface{} `json:"questionId"`
			QuestionNum   interface{} `json:"questionNum"`
			Answer        interface{} `json:"answer"`
			Input         interface{} `json:"input"`
			Level         int         `json:"level"`
			Cet           int         `json:"cet"`
			Right         interface{} `json:"right"`
		}{
			{
				PaperDetailId: "us7C9sNKOFy8unEe202",
				Title:         "culture . ",
				AnswerA:       "时尚 . ",
				AnswerB:       "文化 . ",
				AnswerC:       "狂热 . ",
				AnswerD:       "品种 . ",
			},
		},
	}
	ans, err := GetAns(ques)
	if err != nil {
		t.Error(err)
	}
	t.Log(ans)
}
