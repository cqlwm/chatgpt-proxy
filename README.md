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

docker run --name chatgpt-proxy-v4-c -e ACCESS_TOKEN="eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJseW5xenZlMDY4NjQzQDE2My5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZX0sImh0dHBzOi8vYXBpLm9wZW5haS5jb20vYXV0aCI6eyJ1c2VyX2lkIjoidXNlci1pR3dWYjRTcGd1bDRzSTZ1NVZRazF4clgifSwiaXNzIjoiaHR0cHM6Ly9hdXRoMC5vcGVuYWkuY29tLyIsInN1YiI6ImF1dGgwfDYzYTg0MWZhZjE0NGQ2NjU3NWUzMzY2MiIsImF1ZCI6WyJodHRwczovL2FwaS5vcGVuYWkuY29tL3YxIiwiaHR0cHM6Ly9vcGVuYWkub3BlbmFpLmF1dGgwYXBwLmNvbS91c2VyaW5mbyJdLCJpYXQiOjE2Nzg0NjU2MDUsImV4cCI6MTY3OTY3NTIwNSwiYXpwIjoiVGRKSWNiZTE2V29USHROOTVueXl3aDVFNHlPbzZJdEciLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIG1vZGVsLnJlYWQgbW9kZWwucmVxdWVzdCBvcmdhbml6YXRpb24ucmVhZCBvZmZsaW5lX2FjY2VzcyJ9.AWgWCnCZFyBxABGb63NkeuT0CbYrEedehcXNwAHLBOPHeRsf0kPbB0CkPT_40Z45HZ047z5wbeuYOeEN9bWj_LJElfAE5oNPBotPu6z0geHQITalUKaz6VOX5he_aF74g1Nfjxk7a4jpc8KP22fnvaMCK-Ycln_tdq7WhJMc8FVwn4nI_aASMATxZvc1pPnerr-_c-f4Upe-urmZY8IDMPjJqEpNH3NUwiiu7hLb2e5bico6Ku3Dv_htIeWovtMK18dVbjAbBP6Q1YsTs-bU3C1rwCzqJM_2PQUIocYFbWn9mSUgLjyl2Kr-kdOtRSXEa38uoYGeNAbzGy4UMz_78w" -e PUID="user-iGwVb4Spgul4sI6u5VQk1xrX"  -p 8789:80 -d chatgpt-proxy-v4:0.0.1
```





















