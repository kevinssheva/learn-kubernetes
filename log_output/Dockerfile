# Use a builder image to compile the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod ./
# RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /log-output .

# Use a minimal base image for the final container
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /log-output /log-output

# Set the entrypoint for the container
ENTRYPOINT ["/log-output"]
