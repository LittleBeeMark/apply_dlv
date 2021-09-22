FROM golang:1.17.1-alpine3.14

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl

RUN  GOOS=linux GOARCH=amd64 go install github.com/go-delve/delve/cmd/dlv@latest
RUN  ls /go/bin/linux_amd64 && pwd && go env && \
  cp $GOPATH/bin/linux_amd64/dlv /usr/local/bin/dlv

COPY hello /usr/local/hello
WORKDIR /usr/local/


CMD dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./hello

