# Set the base directory name
BASEDIR := {AppName}

# Default target when 'make' is run without arguments
.DEFAULT_GOAL := build

# Install dependencies
install:
	go mod tidy

# Download and install air for hot reloading
setup:
	@if ! [ -x "$(go env GOPATH)/bin/air" ]; then \
		echo "Installing air..."; \
		curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin; \
	else \
		echo "Air is already installed"; \
	fi

# Build the Go project
build:
	go build -o $(BASEDIR)

# Run the Go project
run: build
	./$(BASEDIR)

# Clean up the build artifacts
clean:
	rm -f $(BASEDIR)

# Run the application in development mode with hot reloading
dev: setup
	$(go env GOPATH)/bin/air
