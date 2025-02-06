# Use Go base image
FROM golang:1.23.4

# Set working directory
WORKDIR /app

# Copy source code
COPY . .
RUN go mod download



# Build the gRPC server
RUN go build -o grpc-server .
#------------


# Expose gRPC port
EXPOSE 50051

# Run the gRPC server
CMD ["./grpc-server"]
