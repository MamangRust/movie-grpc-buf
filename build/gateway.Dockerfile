FROM golang:1.22 AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o gateway ./cmd/gateway

FROM alpine
# Create a directory and copy the binary with correct permissions
WORKDIR /app
COPY --from=builder /build/gateway .


# Use the full path to the binary
CMD ["/app/gateway"]