FROM golang:1.11-alpine
MAINTAINER Jordan Knott <jordan@jordanthedev.com>

RUN apk update && \
    apk add git && \
    rm -r /var/cache/apk/*

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
