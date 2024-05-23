package api

import (
	"Iltmw/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	apiKey      = "xcqrCYyJ76uSFqeHgb3i0IDw"
	secretKey   = "renAUjjDoeTHjok6ViPS06ClP5yIIj4w"
	apiEndpoint = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ernie_speed"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Messages []Message `json:"messages"`
}

// 假设的 API 响应结构体
type ResponseBody struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Created          int64  `json:"created"`
	Result           string `json:"result"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
	Usage            struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}
}

//func ConvertToText(ques *model.ModelObj) string {
//	res := make([]model.Question, 100)
//	for i, _ := range ques.List {
//		res1 := model.Question{
//			PaperDetailId: ques.List[i].PaperDetailId,
//			Title:         ques.List[i].Title,
//			AnswerA:       ques.List[i].AnswerA,
//			AnswerB:       ques.List[i].AnswerB,
//			AnswerC:       ques.List[i].AnswerC,
//			AnswerD:       ques.List[i].AnswerD,
//		}
//		res = append(res, res1)
//	}
//	m, _ := json.Marshal(res)
//	return string(m)
//}

//func ConvertToText(ques *model.ModelObj) string {
//	// 转换 ModelObj 为文本格式
//	// 假设 ModelObj 包含一个 List 字段，其每个元素都有 Title 和四个答案选项
//	var text string
//	for _, item := range ques.List {
//		text += item.Title + "\nA. " + item.AnswerA + "\nB. " + item.AnswerB + "\nC. " + item.AnswerC + "\nD. " + item.AnswerD + "\n"
//	}
//	return text
//}

//func GetAns(ques *model.ModelObj) (*model.Result, error) {
//	tokenResp, err := GetToken(apiKey, secretKey)
//	if err != nil {
//		fmt.Println("take token is error")
//		return nil, err
//	}
//
//	text := ConvertToText(ques)
//	messages := []Message{
//		{Role: "user", Content: "我会给你5题翻译题，你只需要给我返回正确的选项，你需要做到百分之八十的正确率，题目如下：" + text},
//	}
//
//	// 确保消息数量为奇数
//	if len(messages)%2 == 0 {
//		messages = append(messages, Message{Role: "system", Content: "确保消息数量为奇数"})
//	}
//
//	reqBody := RequestBody{
//		Messages: messages,
//	}
//
//	u, err := url.Parse(apiEndpoint)
//	if err != nil {
//		fmt.Println("error parsing url ")
//		return nil, fmt.Errorf("error parsing URL: %w", err)
//	}
//
//	query := u.Query()
//	query.Set("access_token", tokenResp.AccessToken)
//	u.RawQuery = query.Encode()
//
//	body, err := json.Marshal(reqBody)
//	if err != nil {
//		fmt.Println("error marshalling request body ")
//		return nil, fmt.Errorf("error marshalling request body: %w", err)
//	}
//
//	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
//	if err != nil {
//		return nil, fmt.Errorf("error creating request: %w", err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("error sending request ")
//		return nil, fmt.Errorf("error sending request: %w", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("failed to get answer: status code %d", resp.StatusCode)
//	}
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, fmt.Errorf("error reading response body: %w", err)
//	}
//	fmt.Println("respBody: ", string(respBody))
//	var response ResponseBody
//	err = json.Unmarshal(respBody, &response)
//	if err != nil {
//		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
//	}
//	fmt.Println("response: ", response.Result)
//
//	var res model.Result
//	res.PaperId = ques.PaperId
//	res.Type = ques.Type
//	// ------------------------------
//	var cnt int
//	for i := 0; i < len(response.Result); i++ {
//		result := string(response.Result[i])
//		if result == "A" || result == "B" || result == "C" || result == "D" {
//			res.List = append(res.List, model.Answer{
//				Input:         result,
//				PaperDetailId: ques.List[cnt].PaperDetailId,
//			})
//			cnt++
//		}
//	}
//	//---------------------------
//
//	return &res, nil
//
//}

func ConvertToText(ques *model.ModelObj, start, end int) string {
	var text string
	for _, item := range ques.List[start:end] {
		text += item.Title + "\nA. " + item.AnswerA + "\nB. " + item.AnswerB + "\nC. " + item.AnswerC + "\nD. " + item.AnswerD + "\n"
	}
	return text
}

func GetAns(ques *model.ModelObj) (*model.Result, error) {
	tokenResp, err := GetToken(apiKey, secretKey)
	if err != nil {
		fmt.Println("take token is error")
		return nil, err
	}

	var allAnswers = make([]model.Answer, 99)
	var res model.Result
	var cnt int
	res.List = make([]model.Answer, 100)
	batchSize := 5
	for i := 0; i < len(ques.List); i += batchSize {
		end := i + batchSize
		if end > len(ques.List) {
			end = len(ques.List)
		}

		text := ConvertToText(ques, i, end)
		messages := []Message{
			{Role: "user", Content: "我会给你5题翻译题，你一定要给我返回正确答案的序号，一道翻译题中只能出现一个大写字母即正确答案的序号，返回的答案的格式为`x-x-x-x-x`，x表示正确的序号，题目如下：" + text},
		}

		// 确保消息数量为奇数
		if len(messages)%2 == 0 {
			messages = append(messages, Message{Role: "system", Content: "确保消息数量为奇数"})
		}

		reqBody := RequestBody{
			Messages: messages,
		}

		u, err := url.Parse(apiEndpoint)
		if err != nil {
			fmt.Println("error parsing url ")
			return nil, fmt.Errorf("error parsing URL: %w", err)
		}

		query := u.Query()
		query.Set("access_token", tokenResp.AccessToken)
		u.RawQuery = query.Encode()

		body, err := json.Marshal(reqBody)
		if err != nil {
			fmt.Println("error marshalling request body ")
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}

		req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: 30 * time.Second} // 设置超时
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error sending request ")
			return nil, fmt.Errorf("error sending request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get answer: status code %d", resp.StatusCode)
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		fmt.Println("respBody: ", string(respBody))
		var response ResponseBody
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling response body: %w", err)
		}

		fmt.Println("response: ", response.Result)

		reStr := Rematch(response.Result)

		split := strings.Split(reStr, "-")
		for k := 1; k <= 5; k++ {
			if k <= len(split) {
				s := split[k]
				if s == "A" || s == "B" || s == "C" || s == "D" {
					allAnswers[cnt] = model.Answer{
						Input:         s,
						PaperDetailId: ques.List[cnt].PaperDetailId,
					}
					cnt++
				}
			} else {
				allAnswers[cnt] = model.Answer{
					Input:         "C",
					PaperDetailId: ques.List[cnt].PaperDetailId,
				}
				cnt++
			}
		}
	}
	res.PaperId = ques.PaperId
	res.Type = ques.Type
	res.List = allAnswers

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("res: ", res)
	return &res, nil
}

func Rematch(resp string) string {
	pattern := `[A-D]-[A-D]-[A-D]-[A-D]-[A-D]`

	re := regexp.MustCompile(pattern)

	matches := re.FindString(resp)

	//for _, match := range matches {
	//	fmt.Println("match: ", match)
	//}
	return matches
}
