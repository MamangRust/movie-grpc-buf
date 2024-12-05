# Stage 1: Build
FROM golang:1.22-alpine AS builder

# Install dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory
WORKDIR /build

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -o /app ./cmd/app

# Stage 2: Run
FROM alpine:latest

# Install certificates for HTTPS and curl for downloading wait-for-it
RUN apk --no-cache add bash ca-certificates curl

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app /root/app
RUN chmod +x /root/app

# Download wait-for-it.sh and make it executable
RUN curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x /usr/local/bin/wait-for-it.sh

# Define environment variables for database host and port
ENV DB_HOST=db
ENV DB_PORT=5432

# Run the application with wait-for-it
CMD ["/root/app"]
