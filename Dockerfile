FROM golang:alpine
RUN apk --update add curl git && rm -rf /var/cache/apk/*
RUN mkdir /go/src/work
WORKDIR /go/src/work
