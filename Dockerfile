# Stage 1: Build the Go binary
FROM golang:latest AS builder

# Set environment variables for Go
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o guess-game main.go

# Stage 2: Minimal runtime container
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/guess-game .

# Command to run the executable
CMD ["./guess-game"]
