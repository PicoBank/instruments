FROM golang:1.10-alpine3.7

# Install packages
RUN apk add --no-cache \
        build-base \
        gcc \
        wget \
        curl \
        git \
        bash

# Install DevTools
RUN go get -u github.com/beego/bee; \
    go get -u github.com/pressly/goose/cmd/goose; \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

EXPOSE 8080
