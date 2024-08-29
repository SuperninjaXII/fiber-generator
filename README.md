# Fiber Generator

`fiber-generator` is a CLI tool that helps you quickly set up a basic Go Fiber project. It creates the necessary directories and files, initializes a Go module, and sets up templates for views, public assets, routes, and controllers.

## Features

- Automatically generates a basic project structure for a Go Fiber app.
- Creates directories and files for views, public assets, routes, and controllers.
- Replaces placeholders in template files with the project name.
- Initializes a Go module using `go mod init`.
- Includes a `Makefile` to manage common tasks like installing dependencies, building, and running the project.

## Requirements

- Go (version 1.16 or later)
- Make (optional, but recommended for using the `Makefile`)

## Installation

# using the `go install` command 

```bash
go install github.com/SuperninjaXII/Fiber-generator@latest
```
# using the binaries in the release page

## Usage
the `fiber-gen` cli app accepts a `--name` flag . the name provided will be used in the `go mod init`
and make file
```bash
fiber-generator --name my-project
```
