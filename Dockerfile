FROM golang:1.17-alpine

ENV CGO_ENABLED=0

RUN mkdir /app
WORKDIR /app
CMD go run main.go
EXPOSE 8080