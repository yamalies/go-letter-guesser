FROM golang:latest AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o guess-game main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/guess-game .
CMD ["./guess-game"]
