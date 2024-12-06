FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN apt-get update && apt-get install -y postgresql-client

EXPOSE 8080

RUN go build -o /app/main ./cmd

ENV DOCKER=true

CMD ["/app/main"]
