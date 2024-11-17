FROM golang:latest AS build
WORKDIR /baseball_api
COPY . .
RUN go mod download
WORKDIR /baseball_api/cmd/baseball_api
RUN go build -o main

FROM golang:latest
WORKDIR /baseball_api
COPY --from=build /baseball_api/cmd/baseball_api/main ./main
COPY --from=build /baseball_api/assets ./assets
EXPOSE 4242
CMD ["./main"]
