FROM golang:1.16-alpine AS build
WORKDIR /build
ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .
RUN go mod vendor
COPY src ./src
