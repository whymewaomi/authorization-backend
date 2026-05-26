# Dockerfile
FROM golang:1.26.2
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

RUN ls -la /app/main

CMD ["./main"]