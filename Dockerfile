FROM golang:latest

WORKDIR /usr/src/baseball_api

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY .env ./
COPY *.csv ./
COPY *.go ./

RUN go build -o main

EXPOSE 4242
