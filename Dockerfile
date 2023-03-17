FROM golang:1.18-alpine as build

LABEL Maintainer="Yunus Eskisan <yunus.eskisan@hotmail.co.uk>"

ARG SERVICE

RUN mkdir src/emre

WORKDIR /src/emre

COPY . .

RUN go build ${SERVICE}/main.go

ENTRYPOINT ["./main"]
