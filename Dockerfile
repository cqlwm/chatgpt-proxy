FROM golang:1.18-alpine as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app .

RUN mkdir build && cp app build

FROM alpine

WORKDIR /app

COPY --from=builder /app/build .

ENV GIN_MODE=release \
    PORT=80

EXPOSE 80

ENTRYPOINT ["./app"]