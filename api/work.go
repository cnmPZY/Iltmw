package api

import (
	"Iltmw/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	apiKey      = ""
	secretKey   = ""
	apiEndpoint = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_preemptible"
)

// 假设的 API 请求结构体
type RequestBody struct {
	Text string `json:"text"` // 你要发送给文心一言的文本
}

// 假设的 API 响应结构体
type ResponseBody struct {
	Answer string `json:"answer"` // 文心一言的响应文本
}

func ConvertToText(ques *model.ModelObj) string {
	res := make([]model.Question, 100)
	for i, _ := range ques.List {
		res1 := model.Question{
			PaperDetailId: ques.List[i].PaperDetailId,
			Title:         ques.List[i].Title,
			AnswerA:       ques.List[i].AnswerA,
			AnswerB:       ques.List[i].AnswerB,
			AnswerC:       ques.List[i].AnswerC,
			AnswerD:       ques.List[i].AnswerD,
		}
		res = append(res, res1)
	}
	m, _ := json.Marshal(res)
	return string(m)
}

func GetAns(ques *model.ModelObj) (*model.Result, error) {
	// 替换为你的 API 密钥和端点

	tokenResp, err := GetToken(apiKey, secretKey)
	if err != nil {
		return nil, err
	}
	text := ConvertToText(ques)
	reqBody := RequestBody{
		Text: "我会给你一串翻译题，返回给我正确的序号，例如A,B,C,D.题目如下：" + text,
	}

	u, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	query := u.Query()
	query.Set("access_token", tokenResp.AccessToken)
	u.RawQuery = query.Encode()

	body, err := json.Marshal(reqBody)
	if err != nil {
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

	var res model.Result
	res.PaperId = ques.PaperId
	res.Type = ques.Type
	for i, _ := range response.Answer {
		list1 := model.Answer{
			Input:         strconv.Itoa(int(response.Answer[i])),
			PaperDetailId: ques.List[i].PaperDetailId,
		}
		res.List = append(res.List, list1)
	}
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	fmt.Println("Answer:", response.Answer)
	return &model.Result{}, nil
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
