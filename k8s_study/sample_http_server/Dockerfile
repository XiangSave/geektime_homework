ARG APP_NAME=httpServer

FROM golang:1.15 as builder
LABEL stage=builder

ARG APP_NAME
RUN GOPROXY="https://goproxy.cn,direct"

COPY . /go/src/app/

WORKDIR /go/src/app/cmd/${APP_NAME}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .




FROM alpine:3.14
#FROM ubuntu as prod
ARG APP_NAME

WORKDIR /root
COPY --from=builder /go/src/app/cmd/${APP_NAME} .
EXPOSE 80
ENTRYPOINT ["./httpServer"]
