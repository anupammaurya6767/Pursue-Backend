# Start from the official Go image
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go get && \
    go build -o backend .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the binary
CMD ["./backend"]
