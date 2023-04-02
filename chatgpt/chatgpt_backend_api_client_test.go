package chatgpt

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func t2() {
	query := url.Values{}
	query.Set("offset", "0")
	query.Set("limit", "20")

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+Token)
	headers.Set("Cookie", Cookie)
	headers.Set("Re-Domain", "https://chat.openai.com")
	headers.Set("Hf-Access-Key", "UgFUrwVGktW9XbkozneV")
	headers.Set("authority", "chat.openai.com")
	headers.Set("accept", "*/*")
	headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	headers.Set("content-type", "application/json")
	headers.Set("referer", "https://chat.openai.com/chat?model=text-davinci-002-render-sha")
	headers.Set("sec-ch-ua", `"Microsoft Edge";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	headers.Set("sec-ch-ua-mobile", "?0")
	headers.Set("sec-ch-ua-platform", `"macOS"`)
	headers.Set("sec-fetch-dest", "empty")
	headers.Set("sec-fetch-mode", "cors")
	headers.Set("sec-fetch-site", "same-origin")
	headers.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.44")

	url := "http://43.134.168.122:8089/backend-api/conversations?" + query.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	defer resp.Body.Close()

	// 在这里读取响应数据
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
	}
	println(string(r))
}
