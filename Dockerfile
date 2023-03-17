FROM golang:1.19.0

WORKDIR /usr/src/beeus-api

RUN go install github.com/cosmtrek/air@latest

COPY . /usr/src/beeus-api

EXPOSE 4444

RUN go mod tidy

RUN go mod vendor

RUN go run ./src/main.go