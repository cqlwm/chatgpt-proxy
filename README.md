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

docker run --name chatgpt-proxy-v4-c -e ACCESS_TOKEN="abc" -e PUID="def"  -p 8789:80 -d chatgpt-proxy-v4:0.0.1
```