package chatgpt

import (
	"net/http"
	"os"
)

const ApplicationJson = "application/json"

var ProxyHost = "http://43.134.168.122:8089"

func init() {
	p := os.Getenv("PROXY_HOST")
	if p != "" {
		ProxyHost = p
	}
}

const Token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJvbmVtb214c2ZrQHlpbmhlbHguY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWV9LCJodHRwczovL2FwaS5vcGVuYWkuY29tL2F1dGgiOnsidXNlcl9pZCI6InVzZXItVklWZmU0TkpXVExyRXdxeHliZWUxa2R3In0sImlzcyI6Imh0dHBzOi8vYXV0aDAub3BlbmFpLmNvbS8iLCJzdWIiOiJhdXRoMHw2NDEzYmUzMjg5Y2U5NWQ3YzYxYTZjNDIiLCJhdWQiOlsiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS92MSIsImh0dHBzOi8vb3BlbmFpLm9wZW5haS5hdXRoMGFwcC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjc5MjA1NzgzLCJleHAiOjE2ODA0MTUzODMsImF6cCI6IlRkSkljYmUxNldvVEh0Tjk1bnl5d2g1RTR5T282SXRHIiwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBtb2RlbC5yZWFkIG1vZGVsLnJlcXVlc3Qgb3JnYW5pemF0aW9uLnJlYWQgb2ZmbGluZV9hY2Nlc3MifQ.qWILc7SyeVGZiq0vX-GG0-0x-Zsel9up0DzoF3NOGRDQ7dt5xoepDRcmO4UqoYMXSmNcHn1c_CQUtzjeYw4l2wvdko9TgsMCS1ikWvNiJO_XSazUS7aDYEIEgUurmykG9B69IaA6S4lqfG1_EAFexuwWi61_kLkMi1Q0zCvtsQFNJDNZUjqLUXlNHhoxGl_LlHQq05-RHGvb7QiilULpd2qqvk3tT2oLfbmE1xfzJX6xYFN65n1q6Svq3IMOpoffnFs9BCeOhL-zAEVy_unt7sqAJncLUUvqVOfEe6raz5X3ZjTHj2vGck6CGrqU9Iq1tzxOsIw1AkqoR_pgXhO3_w`
const Cookie = `_puid=user-VIVfe4NJWTLrEwqxybee1kdw:1679317684-%2B5olaIxL6x3jT8DaFu4fsnLIqnoFBYIEAZmhaTr7LDY%3D`

func UniversalHeader(header *http.Header, cookie string, token string, contentType string) {
	if contentType == "" {
		contentType = ApplicationJson
	}
	header.Set("Authorization", "Bearer "+token)
	header.Set("Cookie", cookie)

	header.Set("Re-Domain", "https://chat.openai.com")
	header.Set("Hf-Access-Key", "UgFUrwVGktW9XbkozneV")
	header.Set("authority", "chat.openai.com")
	header.Set("accept", "*/*")
	header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	header.Set("content-type", "application/json")
	header.Set("referer", "https://chat.openai.com/chat?model=text-davinci-002-render-sha")
	header.Set("sec-ch-ua", `"Microsoft Edge";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	header.Set("sec-ch-ua-mobile", "?0")
	header.Set("sec-ch-ua-platform", `"macOS"`)
	header.Set("sec-fetch-dest", "empty")
	header.Set("sec-fetch-mode", "cors")
	header.Set("sec-fetch-site", "same-origin")
	header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.44")
}
