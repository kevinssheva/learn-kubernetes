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
RUN go build -o /todo-app .

# Use a minimal base image for the final container
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /todo-app /todo-app

# Expose the port
EXPOSE 8080

# Set the entrypoint for the container
ENTRYPOINT ["/todo-app"]
