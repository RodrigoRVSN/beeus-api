# Builder stage
FROM golang:alpine AS builder

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
