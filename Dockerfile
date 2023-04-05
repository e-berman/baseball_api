FROM golang:latest

WORKDIR /usr/src/baseball_api

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY .env ./
COPY *.go ./

RUN go build -o /build/baseball_api

EXPOSE 4242
