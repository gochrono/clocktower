FROM golang:1.12.5-alpine
MAINTAINER Jordan Knott <jordan@jordanthedev.com>

ARG app_env
ENV APP_ENV $app_env

ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on

RUN apk update && \
    apk add git curl && \
    rm -r /var/cache/apk/*

RUN mkdir /go/src/github.com/gochrono/castle -p
WORKDIR /go/src/github.com/gochrono/castle

COPY go.mod .
COPY go.sum .


RUN go get github.com/cespare/reflex

RUN go mod download

COPY . .

ENTRYPOINT ["reflex", "-c", "reflex.conf"]
EXPOSE 8000
