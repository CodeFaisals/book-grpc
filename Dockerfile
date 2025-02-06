# Use Go base image
FROM golang:1.23.4 AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY . .
RUN go mod download



# Build the gRPC server
RUN go build -o grpc-server .
FROM registry.trendyol.com/platform/base/image/appsec/chainguard/static/library:lib-20230201
COPY --from=builder /app/grpc-server /grpc-server
# Expose gRPC port
EXPOSE 50051

# Run the gRPC server
CMD ["./grpc-server"]
