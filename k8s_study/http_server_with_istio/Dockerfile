ARG APP_NAME={{ VAR_SERVER_NAME }}

FROM golang:1.15 as builder

ARG APP_NAME

COPY . /go/src/app/

WORKDIR /go/src/app/cmd/${APP_NAME}
RUN GOPROXY="https://goproxy.cn,direct" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .




FROM alpine:3.14
ARG APP_NAME

WORKDIR /root
COPY --from=builder /go/src/app/cmd/${APP_NAME} .

EXPOSE 8080
VOLUME /root/configs
VOLUME /root/logs


ENTRYPOINT ["./{{ VAR_SERVER_NAME }}"]
