FROM golang:1.13-alpine

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /src

COPY . .

RUN go mod download
