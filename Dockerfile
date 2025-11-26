# Use official Go image as base
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o expense-api .

# Final stage - use minimal alpine image
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/expense-api .

# Expose port
EXPOSE 8080

# Set environment to release mode
ENV GIN_MODE=release

# Run the application
CMD ["./expense-api"]
