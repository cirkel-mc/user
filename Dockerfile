FROM golang:1.21.1-alpine3.18 AS builder

RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR /usr/app

ARG user

ARG token

RUN git config --global url."https://${user}:${token}@github.com".insteadOf "https://github.com"

ENV GOSUMDB=off

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o bin main.go

FROM alpine:3.18

RUN apk add --no-cache tzdata ca-certificates libc6-compat

COPY --from=builder /usr/app/bin /usr/app/bin

ENTRYPOINT ["/usr/app/bin"]