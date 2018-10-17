FROM golang:alpine AS builder
ADD . /src
WORKDIR /src/server
RUN go build -o server
ENTRYPOINT ./server
