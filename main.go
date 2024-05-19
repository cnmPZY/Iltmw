package main

import (
	"Iltmw/sso"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// I love to memorize words
// 这是我爱记单词的脚本

// 智慧杭电登录，拿到一个token
// token 去上课啦里面获取试卷
// json 里面有所有的题目
// 从题目里面提取出所有的单词，搜索到单词的意思
// 匹配完之后呢，你自己去生成一个新的Jason，然后这个新的Jason里面是要有你所有的什么试卷什么东西。
// 再调用提交的接口，把这个新的Json提交上去，就可以了。
// 如果是电脑上做题目，就需要带上手机的user

// https://github.com/hdulib/hdu

const url = "https://sso.hdu.edu.cn/login?service=https:%2F%2Fi.hdu.edu.cn%2Fsopcb%2F"
const username = "23050118"
const password = "E:13819517722@163.com"

func main() {
	login, err := sso.Login(url, username, password)
	if err != nil {
		panic(err)
	}
	token := strings.Split(login[0].String(), "=")
	fmt.Println(token[1])
	//tokens := "d1b2e099-ab7a-4d80-a384-50fc6c793ce7"
	//request(tokens, 11, "0", 0)
}

// PostHeader F12 userinfo?type=5&index=
// PostHeader 智慧杭电的请求头
func PostHeader(token string) http.Header {
	headers := http.Header{}
	headers.Set("Host", "skl.hdu.edu.cn")
	headers.Set("Sec-Ch-Ua", `"Chromium";v="117", "Not;A=Brand";v="8"`)
	headers.Set("Skl-Ticket", generateTicket(21))
	headers.Set("Sec-Ch-Ua-Mobile", "?0")
	headers.Set("User-Agent", "Mozilla/5.0 (Linux; Android 4.2.1; M040 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36")
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json, text/plain, */*")
	headers.Set("X-Auth-Token", token)
	headers.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	headers.Set("Origin", "https://skl.hduhelp.com")
	headers.Set("Sec-Fetch-Site", "cross-site")
	headers.Set("Sec-Fetch-Mode", "cors")
	headers.Set("Sec-Fetch-Dest", "empty")
	headers.Set("Referer", "https://skl.hduhelp.com/")
	headers.Set("Accept-Encoding", "gzip, deflate, br")
	headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	return headers
}

// generateTicket 生成一个随机的 ticket
func generateTicket(length int) string {
	const NL = "useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	ticket := make([]byte, length)
	for i, b := range bytes {
		ticket[i] = NL[b&63] // 与 `& 63` 相同，确保索引在 0-63 范围内
	}
	return string(ticket)
}

// 我爱记单词的header
func GetHeaders(token string) http.Header {
	ticket := generateTicket(21) // 自定义函数，模拟 JavaScript 中的 ticket 函数
	headers := http.Header{}
	headers.Set("Skl-Ticket", ticket)
	headers.Set("X-Auth-Token", token)
	headers.Set("User-Agent", "Mozilla/5.0 (Linux; Android 4.2.1; M040 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36")
	headers.Set("Accept", "application/json, text/plain, */*")
	headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	headers.Set("Connection", "keep-alive")
	headers.Set("Referer", "https://skl.hdu.edu.cn/")
	return headers
}

func request(token string, week int, mode string, delay int) {
	client := http.Client{}
	startTime := time.Now().UnixMilli()
	urls := fmt.Sprintf("https://skl.hdu.edu.cn/api/paper/new?type=%s&week=%d&startTime=%d", mode, week, startTime)
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		panic(err)
	}
	req.Header = GetHeaders(token)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	filename := fmt.Sprintf("paper_%s.json", time.Now().Format("20060102150405"))
	body, _ := io.ReadAll(resp.Body)
	_ = os.WriteFile(filename, body, 0644)
	fmt.Println("存储试卷信息中")

	q := make([]Paper, 100)
	if err := json.Unmarshal(body, &q); err != nil {
		panic(err)
	}

	fmt.Println("等待提交试卷")
	//go progress(delay)
	//for i, v := range q {
	//
	//}
	time.Sleep(time.Duration(delay) * time.Second)

}
