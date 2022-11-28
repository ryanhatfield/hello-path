# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY server.go ./

RUN go build -o=/hello-path .

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /hello-path /hello-path

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/hello-path"]