# syntax=docker/dockerfile:1

# Build step
FROM golang:latest as build

WORKDIR /app

RUN apt-get update
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o build/ant.dev/main main.go 

# Entrypoint step
FROM gcr.io/distroless/base

COPY --from=build /app/. /

ENTRYPOINT ["./build/ant.dev/main"]