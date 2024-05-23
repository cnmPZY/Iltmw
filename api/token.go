package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken(apiKey string, secretKey string) (*TokenResponse, error) {
	baseUrl := "https://aip.baidubce.com/oauth/2.0/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", apiKey)
	data.Set("client_secret", secretKey)

	client := &http.Client{Timeout: 10 * time.Second} // 设置超时

	// 增加重试机制
	var resp *http.Response
	var err error
	for retries := 3; retries > 0; retries-- {
		resp, err = client.PostForm(baseUrl, data)
		if err == nil {
			break
		}
		fmt.Printf("Error sending request, retries left: %d, error: %v\n", retries-1, err)
		time.Sleep(2 * time.Second) // 等待一段时间后重试
	}
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("failed to get token: status code %d, body: %s", resp.StatusCode, bodyString)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return &tokenResp, nil
}
