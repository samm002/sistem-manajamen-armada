# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates and postgresql-client for migrations
RUN apk --no-cache add ca-certificates tzdata postgresql-client

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy migration files
COPY --from=builder /app/database/migration ./migrations

# Copy migrate binary
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy startup script
COPY start.sh .
RUN chmod +x start.sh

# Expose port
EXPOSE 3000

# Use startup script instead of direct command
CMD ["./start.sh"]