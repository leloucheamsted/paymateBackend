FROM golang:1.17-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download
RUN go build -o ./app 

EXPOSE 8080

ENTRYPOINT ["./app"]