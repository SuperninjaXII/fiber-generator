# Fiber Generator

`fiber-generator` is a lightweight CLI tool designed to bootstrap Go Fiber applications with a clean, simple project structure. It automatically sets up directories and boilerplate code, allowing developers to start building their web applications immediately.

## Overview

Fiber Generator creates a straightforward project structure that integrates the high-performance Fiber web framework with HTMX for modern, server-side web applications. The generated project includes routing, controllers, and static file serving out of the box.

## Features

- **Automated Project Scaffolding**
  - Creates a clean directory structure for views, public assets, and controllers
  - Generates a basic Fiber application setup
  - Includes HTMX integration for dynamic web interfaces

- **Development Ready**
  - Initializes a Go module with required dependencies
  - Provides pre-configured static file serving
  - Sets up basic routing structure

## Requirements

- Go 1.20 or later
- Make (optional, recommended for development workflow)

## Installation

Install the CLI tool using Go's package manager:

```bash
go install github.com/SuperninjaXII/fiber-generator@latest
```

## Usage

### Creating a New Project

Generate a new project with:

```bash
fiber-generator --name my-project
```

### Project Structure

The generated project follows this structure:

```
my-project/
├── app.go                 # Main application entry point
├── controllers/
│   └── CreateUserHandler.go
├── routes/
│   └── userRoutes.go
├── views/
│   └── index.html
├── public/
│   ├── css/
│   │   └── style.css
│   ├── js/
│   │   └── app.js
│   └── lib/
│       └── htmx.min.js
├── go.mod
└── Makefile
```

### Setup and Development

1. Navigate to your project directory:
   ```bash
   cd my-project
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the development server:
   ```bash
   make dev
   ```

### Available Make Commands

- `make dev`: Start the development server
- `make setup`: Install project dependencies

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.