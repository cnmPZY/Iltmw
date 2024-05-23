package api

import (
	"Iltmw/model"
	"fmt"
	"testing"
)

func TestGetAns(t *testing.T) {
	ques := &model.ModelObj{
		PaperId: "OS6hpoUIRVM2BLDxVkX",
		Type:    "5",
		List: []model.ListItem{
			{
				PaperDetailId: "us7C9sNKOFy8unEe202",
				Title:         "culture",
				AnswerA:       "时尚",
				AnswerB:       "文化",
				AnswerC:       "狂热",
				AnswerD:       "品种",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe203",
				Title:         "today",
				AnswerA:       "今天",
				AnswerB:       "明天",
				AnswerC:       "昨天",
				AnswerD:       "明年",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe204",
				Title:         "tomorrow",
				AnswerA:       "今天",
				AnswerB:       "明天",
				AnswerC:       "昨天",
				AnswerD:       "明年",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe205",
				Title:         "beauty",
				AnswerA:       "美丽",
				AnswerB:       "丑陋",
				AnswerC:       "妈妈",
				AnswerD:       "丑",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe206",
				Title:         "charge",
				AnswerA:       "充电",
				AnswerB:       "收费",
				AnswerC:       "充值",
				AnswerD:       "充电器",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe207",
				Title:         "change",
				AnswerA:       "改变",
				AnswerB:       "变化",
				AnswerC:       "更换",
				AnswerD:       "变",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe208",
				Title:         "milk",
				AnswerA:       "牛奶",
				AnswerB:       "奶粉",
				AnswerC:       "奶酪",
				AnswerD:       "奶",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe209",
				Title:         "water",
				AnswerA:       "水",
				AnswerB:       "饮料",
				AnswerC:       "酒",
				AnswerD:       "饮水",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe210",
				Title:         "food",
				AnswerA:       "食物",
				AnswerB:       "饮食",
				AnswerC:       "食品",
				AnswerD:       "食",
			},
			{
				PaperDetailId: "us7C9sNKOFy8unEe211",
				Title:         "drink",
				AnswerA:       "喝",
				AnswerB:       "饮料",
				AnswerC:       "饮品",
				AnswerD:       "饮",
			},
		},
	}
	res, err := GetAns(ques)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(*res)
}

func TestGetToken(t *testing.T) {
	token, err := GetToken(apiKey, secretKey)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token.AccessToken)
	//t.Log(token)
}
