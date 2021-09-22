FROM golang:alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add --no-cache git build-base

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY hello /usr/local/hello
WORKDIR /usr/local


CMD $GOPATH/bin/dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./hello
