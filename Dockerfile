FROM golang as builder
MAINTAINER  ogre222
WORKDIR /root/code
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn/,direct
RUN  go build -v  -o ./bin/main main.go

FROM golang
WORKDIR /root/code
COPY --from=builder /root/code .
EXPOSE 12345
ENTRYPOINT ["./bin/main"]