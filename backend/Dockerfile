# 使用官方的 Golang 镜像
FROM golang:1.21-alpine
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
 # 设置工作目录
WORKDIR /backend
COPY . .
RUN go mod tidy
RUN go build -o server .
 # 暴露服务端口
EXPOSE 8080
 # 运行应用
CMD ["./server"]
