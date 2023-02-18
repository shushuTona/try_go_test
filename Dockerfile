FROM golang:1.20

RUN cd /go/src && mkdir try_go_test

WORKDIR /go/src/try_go_test
