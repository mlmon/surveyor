FROM golang:1.25.1

WORKDIR /app

COPY . .

RUN go test -v -cover -covermode atomic -coverprofile cover.out ./...