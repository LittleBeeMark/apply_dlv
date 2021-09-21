FROM alpine

COPY hello /usr/local/hello
WORKDIR /usr/local/

CMD ["./hello"]