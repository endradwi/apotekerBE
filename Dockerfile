# ---- Build Stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (required for go get if you use private repos)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o app

# ---- Run Stage ----
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Copy the upload directory if needed
COPY --from=builder /app/upload ./upload

# Expose the port your app runs on (change if needed)
EXPOSE 8889

# Command to run the binary
CMD ["./app"] 