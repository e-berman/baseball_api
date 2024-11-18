FROM golang:latest AS build
WORKDIR /baseball_api
COPY . .
RUN go mod download
WORKDIR /baseball_api/cmd/baseball_api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest
WORKDIR /baseball_api
RUN apk add --no-cache curl
COPY --from=build /baseball_api/cmd/baseball_api/main ./main
COPY --from=build /baseball_api/assets ./assets
EXPOSE 4242
CMD ["./main"]
