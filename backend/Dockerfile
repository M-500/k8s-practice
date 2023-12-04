FROM golang:alpine as builder
WORKDIR /go/src/kubeimooc.com/server
COPY . .


RUN go env -w GO111MODULE = ON \
    && go env -w GOPROXY = https://goproxy.cn,direct \
    && go env -w CGO_ENABLED = 0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="18574945291@163.com"

WORKDIR /go/src/wll.com/server
COPY --from=0 /go/src/k8spratice.com/server/config.yaml ./config.yaml
COPY --from=0 /go/src/k8spratice.com/server/.kube/config .kube/config
COPY --from = 0 /go/src/k8spratice.com/server/backend ./

EXPOSE 8080
ENTRYPOINT ./server