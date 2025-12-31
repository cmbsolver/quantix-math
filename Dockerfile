# Stage 1: Build the application
FROM golang:1.25-alpine AS builder

# Install build dependencies if needed (e.g., git, gcc)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary for alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o quantix-math main.go

# Stage 2: Run the application
FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/quantix-math .

# Copy runtime assets required by the application
# Based on your main.go, these are necessary:
COPY --from=builder /app/views ./views
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/settings ./settings

# Expose the port the app listens on (from app.Listen(":3301"))
EXPOSE 3301

# Run the binary
CMD ["./quantix-math"]