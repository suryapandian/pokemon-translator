
FROM golang:1.15-alpine AS builder

ADD . /go/src/pokemon-translator

WORKDIR /go/src/pokemon-translator

RUN go build -mod=vendor -o pokemon-translator .

EXPOSE 5000

ENTRYPOINT [ "./pokemon-translator"]