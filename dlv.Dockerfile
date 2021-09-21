FROM golang:1.17.1-alpine3.14 as build

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /usr/local/bin/dlv github.com/go-delve/delve/cmd/dlv

COPY hello /usr/local/hello
WORKDIR /usr/local/


CMD ["dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./hello"]

