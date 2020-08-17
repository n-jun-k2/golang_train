FROM golang:1.15.0-buster

ENV GOBIN ${GOPATH}/bin
RUN mkdir /tmp/app && go get github.com/lib/pq

WORKDIR /tmp/app
COPY ./src ./