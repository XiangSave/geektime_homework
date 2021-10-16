FROM golang:1.15 as builder
ARG appName=httpServer

LABEL author=xiang
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN cd cmd/${appName} && go build .



FROM alpine:3.14 as prod
WORKDIR /root/
COPY --from=builder /go/src/app/cmd/httpServer/httpServer /root
CMD ["./httpServer"]
