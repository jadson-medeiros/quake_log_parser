# Use the Golang base image in Alpine version
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire source code to the working directory
COPY . .

# Compile the Go code
RUN go build -o main ./cmd/quake_log_parser

# Define the entry point for the container
ENTRYPOINT ["./main"]
