package api

import (
	"Iltmw/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiKey      = "xcqrCYyJ76uSFqeHgb3i0IDw"
	secretKey   = "renAUjjDoeTHjok6ViPS06ClP5yIIj4w"
	apiEndpoint = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/yi_34b_chat"
)

// 假设的 API 请求结构体
type RequestBody struct {
	Text string `json:"text"` // 你要发送给文心一言的文本
}

// 假设的 API 响应结构体
type ResponseBody struct {
	Answer string `json:"answer"` // 文心一言的响应文本
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

func ConvertToText(ques *model.ModelObj) string {
	// 转换 ModelObj 为文本格式
	// 假设 ModelObj 包含一个 List 字段，其每个元素都有 Title 和四个答案选项
	var text string
	for _, item := range ques.List {
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
	text := ConvertToText(ques)
	reqBody := RequestBody{
		Text: "我会给你一串翻译题，返回给我正确的序号，例如A,B,C,D题目如下：" + text,
	}
	var b []byte
	_ = json.Unmarshal(b, &reqBody)
	if len(b)%2 == 1 {
		reqBody.Text += "."
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

	client := &http.Client{}
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

	var res model.Result
	res.PaperId = ques.PaperId
	res.Type = ques.Type
	for i, ans := range response.Answer {
		list1 := model.Answer{
			Input:         ans,
			PaperDetailId: ques.List[i].PaperDetailId,
		}
		res.List = append(res.List, list1)
	}

	fmt.Println("Answer:", response.Answer)
	return &res, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken(apiKey string, secretKey string) (*TokenResponse, error) {
	baseUrl := "https://aip.baidubce.com/oauth/2.0/token"

	u, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Println("Error parsing url...")
		return nil, err
	}

	query := u.Query()
	query.Set("grant_type", "client_credentials")
	query.Set("client_id", apiKey)
	query.Set("client_secret", secretKey)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get the baidu token...")
	}

	var token TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}
