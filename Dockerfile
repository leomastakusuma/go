FROM golang:1.12.0-alpine3.9
RUN apk add -q --update \
    && apk add -q \
            bash \
            git \
            curl \
    && rm -rf /var/cache/apk/*
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql

WORKDIR /go/src/github.com

RUN mkdir second

COPY  . /go/src/github.com/second

WORKDIR /go/src/github.com/second

CMD ["go","run","main.go"]