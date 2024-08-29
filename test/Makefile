# Makefile for the fiber-generator project

# Set the base directory name
BASEDIR := test

# Default target when 'make' is run without arguments
.DEFAULT_GOAL := build

# Install dependencies by running 'go mod tidy'
install:
	go mod tidy

# Build the Go project
build:
	go build -o $(BASEDIR)

# Run the Go project
run: build
	./$(BASEDIR)

# Clean up the build artifacts
clean:
	rm -f $(BASEDIR)