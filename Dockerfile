# FROM golang:latest
FROM golang:latest

# set the working directory inside the container
WORKDIR /baseball_api

# copy contents
COPY . ./

# download dependencies
RUN go mod download

# change working directory to where main.go is stored
WORKDIR /baseball_api/cmd/baseball_api

# build executable
RUN go build -o main

# expose api port
EXPOSE 4242

# run the executable
CMD ["./main"]
