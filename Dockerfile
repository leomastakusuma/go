FROM golang:1.12.0-alpine3.9
RUN mkdir /github.com/second
ADD . /github.com/second
WORKDIR /github.com/second
RUN go get -v
RUN go build -o main .
CMD ["go","run","main.go"]