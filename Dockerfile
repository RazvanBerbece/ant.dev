# syntax=docker/dockerfile:1

# Base Image
FROM golang:latest

WORKDIR /app

RUN apt-get update
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o build/ant.dev/main main.go 

CMD sh -c "./build/ant.dev/main"