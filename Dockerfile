# Build Stage
FROM golang:1.21-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o image-viewer .

# Runtime Stage
FROM alpine:3.18
ENV GIN_MODE=release
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Install dependencies (only needed if you use TLS/HTTPS)

# Create app directory
WORKDIR /app

# Copy binary and templates from builder
COPY --from=builder /app/image-viewer .
COPY --from=builder /app/templates ./templates

# Create image directory
RUN mkdir -p /app/images

# Environment variables
ENV UPLOAD_FOLDER=/app/images
ENV PORT=8080

# Expose port
EXPOSE 8080

# Start the application
CMD ["./image-viewer"]
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/ || exit 1
