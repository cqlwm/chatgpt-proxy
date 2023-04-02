# ChatGPT-Proxy-V4
Cloudflare Bypass for OpenAI based on `puid`

fork project:
https://github.com/acheong08/ChatGPT-Proxy-V4

## Requirements

- ChatGPT plus account 
- Access to chat.openai.com

## Installation 

`go install github.com/acheong08/ChatGPT-Proxy-V4@latest`

## Usage

### Environment variables
- `ACCESS_TOKEN` - For automatic refresh of `_puid`
- `PUID` - Preset `_puid`

Choose one or both.

### Running
`ChatGPT-Proxy-V4`

## start
```shell
docker build -t chatgpt-proxy-v4:0.0.1 .

docker run --name chatgpt-proxy-v4-c -e ACCESS_TOKEN="eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJvbmVtb214c2ZrQHlpbmhlbHguY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWV9LCJodHRwczovL2FwaS5vcGVuYWkuY29tL2F1dGgiOnsidXNlcl9pZCI6InVzZXItVklWZmU0TkpXVExyRXdxeHliZWUxa2R3In0sImlzcyI6Imh0dHBzOi8vYXV0aDAub3BlbmFpLmNvbS8iLCJzdWIiOiJhdXRoMHw2NDEzYmUzMjg5Y2U5NWQ3YzYxYTZjNDIiLCJhdWQiOlsiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS92MSIsImh0dHBzOi8vb3BlbmFpLm9wZW5haS5hdXRoMGFwcC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjc5MjA1NzgzLCJleHAiOjE2ODA0MTUzODMsImF6cCI6IlRkSkljYmUxNldvVEh0Tjk1bnl5d2g1RTR5T282SXRHIiwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBtb2RlbC5yZWFkIG1vZGVsLnJlcXVlc3Qgb3JnYW5pemF0aW9uLnJlYWQgb2ZmbGluZV9hY2Nlc3MifQ.qWILc7SyeVGZiq0vX-GG0-0x-Zsel9up0DzoF3NOGRDQ7dt5xoepDRcmO4UqoYMXSmNcHn1c_CQUtzjeYw4l2wvdko9TgsMCS1ikWvNiJO_XSazUS7aDYEIEgUurmykG9B69IaA6S4lqfG1_EAFexuwWi61_kLkMi1Q0zCvtsQFNJDNZUjqLUXlNHhoxGl_LlHQq05-RHGvb7QiilULpd2qqvk3tT2oLfbmE1xfzJX6xYFN65n1q6Svq3IMOpoffnFs9BCeOhL-zAEVy_unt7sqAJncLUUvqVOfEe6raz5X3ZjTHj2vGck6CGrqU9Iq1tzxOsIw1AkqoR_pgXhO3_w" -e SET_COOKIE="_puid=user-VIVfe4NJWTLrEwqxybee1kdw:1679209981-ajllEf9OrgH%2FN29D60E1MuWCzz7Dd%2FHhpkxrfpGT8BE%3D"  -p 8789:80 -d chatgpt-proxy-v4:0.0.1
```

accessToken = os.Getenv("ACCESS_TOKEN")
setCookie   = os.Getenv("SET_COOKIE")

front -> flow -> push -> prompt
server -> flow -> pull -> answer

server -> flow > pull






















