# Define variables that can be reused
BINARY_NAME=quake_log_parser
DOCKER_IMAGE_NAME=quake_log_parser:latest

.PHONY: all build test lint clean run docker-build docker-run fmt

# The default target is build
all: build

# Compile the project
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) ./cmd/quake_log_parser

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Run the linter
lint:
	@echo "Running linter..."
	golangci-lint run ./...

# Clean generated binary files
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

# Run the application
run: build
	@echo "Running the application..."
	./$(BINARY_NAME)

# Display help
help:
	@echo "Makefile commands:"
	@echo "  all         - Build the project (default)"
	@echo "  build       - Build the project"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linter"
	@echo "  clean       - Clean build artifacts"
	@echo "  run         - Run the application"
	@echo "  docker-build- Build Docker image"
	@echo "  docker-run  - Run Docker container"
	@echo "  fmt         - Format the code"

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run Docker container
docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm -it $(DOCKER_IMAGE_NAME)

# Format the code
fmt:
	@echo "Formatting the code..."
	go fmt ./...
