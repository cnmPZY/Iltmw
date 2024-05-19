package main

type Paper struct {
	PaperId  string     `json:"paperId"`
	Type     int        `json:"type"`
	Question []Question `json:"question"`
}

type Question struct {
	PaperDetailId string  `json:"paperDetailId"`
	Title         string  `json:"title"`
	AnswerA       string  `json:"answerA"`
	AnswerB       string  `json:"answerB"`
	AnswerC       string  `json:"answerC"`
	AnswerD       string  `json:"answerD"`
	QuestionId    *string `json:"questionId"`
	QuestionNum   int     `json:"questionNum"`
	Answer        *string `json:"answer"`
	Input         *string `json:"input"`
	Level         int     `json:"level"`
	Cet           int     `json:"cet"`
	Right         *bool   `json:"right"`
}
