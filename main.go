package main

import (
	api "Iltmw/api"
	"Iltmw/model"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

const tokens = "866fc910-6db3-49d4-8e61-cc7132b23316" // 输入你的 token
const week = 5    // 输入这是第几周
const mode = 0    // 0 为自测，1 为考试

func main() {
	request(tokens, week, strconv.Itoa(mode))
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

func request(token string, week int, mode string) {
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

	filename := fmt.Sprintf("./paper/paper_%s.json", time.Now().Format("20060102150405"))
	body, _ := io.ReadAll(resp.Body)
	_ = os.WriteFile(filename, body, 0644)
	fmt.Println("存储试卷信息中")

	q := new(model.ModelObj)
	if err := json.Unmarshal(body, q); err != nil {
		panic(err)
	}
	fmt.Println("试卷信息存储完毕")
	res, err := api.GetAns(q)
	if err != nil {
		panic("get ans is wrong...")
		return
	}
	fmt.Println("等待提交试卷")

	time.Sleep(5 * time.Second)
	err = Submit(res, token)
	if err != nil {
		panic(err)
	}
}

func Submit(res *model.Result, token string) error {
	time.Sleep(5 * time.Second)
	b, _ := json.Marshal(res)
	req, err := http.NewRequest(http.MethodPost, "https://skl.hdu.edu.cn/api/paper/save", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header = GetHeaders(token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.118 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("提交试卷成功")
	if resp.StatusCode == 201 {
		fmt.Println("你不能短时间内自测/考试")
	}

	fmt.Println("提交试卷后的返回结果为：", resp)

	return nil

}
