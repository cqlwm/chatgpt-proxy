package main

import (
	"io"
	"log"
	"os"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(tls_client.Chrome_110),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}
	client, _   = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	accessToken = os.Getenv("ACCESS_TOKEN")
	setCookie   = os.Getenv("SET_COOKIE")
)

func UniversalHeader(header *http.Header, cookie string) {
	header.Set("authority", "chat.openai.com")
	header.Set("accept", "*/*")
	header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("content-type", "application/json")
	header.Set("cookie", cookie)
	header.Set("referer", "https://chat.openai.com/chat?model=text-davinci-002-render-sha")
	header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.44")
}

func main() {
	if accessToken == "" || setCookie == "" {
		panic("Error: ACCESS_TOKEN and PUID are not set")
	}
	go func() {
		url := "https://chat.openai.com/backend-api/models"
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		UniversalHeader(&req.Header, setCookie)

		for {
			resp, err := client.Do(req)
			if err != nil {
				break
			}

			log.Println("Got response: " + resp.Status)
			if resp.StatusCode != 200 {
				body, _ := io.ReadAll(resp.Body)
				log.Println(string(body))
				break
			}
			_ = resp.Body.Close()

			setCookie = resp.Header.Get("Set-Cookie")
			log.Println("New cookie: " + setCookie)

			// Sleep for 6 hour
			time.Sleep(6 * time.Hour)
		}
		panic("Error: Failed to refresh puid cookie")
	}()

	handler := gin.Default()
	handler.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	handler.Any("/api/*path", proxy)

	err := endless.ListenAndServe(":80", handler)
	if err != nil {
		panic(err.Error())
	}
}

func proxy(c *gin.Context) {

	var url string
	var err error
	var requestMethod string
	var request *http.Request
	var response *http.Response

	url = "https://chat.openai.com/backend-api" + c.Param("path")
	requestMethod = c.Request.Method

	request, err = http.NewRequest(requestMethod, url, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	requestToken := c.Request.Header.Get("Authorization")
	requestCookie := c.Request.Header.Get("Cookie")
	if requestCookie == "" && requestToken == "" {
		requestToken = accessToken
		requestCookie = setCookie
	}
	request.Header.Set("Authorization", requestToken)
	request.Header.Set("Cookie", requestCookie)

	response, err = client.Do(request)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", response.Header.Get("Content-Type"))
	c.Status(response.StatusCode)
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, response.Body)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return false
	})

}
