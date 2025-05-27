# Order Packs Calculator

A web application that calculates the optimal number of packs needed to fulfill orders based on available pack sizes.

## Overview

This application helps businesses optimize their packaging strategy by calculating the most efficient combination of packs needed to fulfill any order size. It provides a simple web interface to manage pack sizes and calculate pack distributions.

## Features

- Manage available pack sizes (add, remove)
- Calculate optimal pack combinations for orders
- RESTful API for integration with other systems
- Simple and intuitive web interface

## Technology Stack

- Backend: Go (Golang) with Gin framework
- Frontend: HTML, JavaScript
- Containerization: Docker
- Deployment: Heroku

## API Endpoints

- `GET /api/packs` - Get all available pack sizes
- `POST /api/packs` - Add a new pack size
- `DELETE /api/packs/{size}` - Remove a pack size
- `POST /api/calculate` - Calculate packs needed for an order size

## Development

### Prerequisites

- Go 1.24 or higher
- Docker (for containerized development)
- Heroku CLI (for deployment)

### Setup and Running Locally

1. Clone the repository
2. Build and run the application:

```bash
make build
make run
```

3. Access the application at http://localhost:8080

### Testing

Run tests using:

```bash
make test
```

Generate test coverage report:

```bash
make test-coverage
```

### Docker Support

Build the Docker image:

```bash
make docker-build
```

Run the Docker container:

```bash
make docker-run
```

### Deployment

Deploy the application to Heroku:

```bash
make heroku-deploy
```

Note for Apple Silicon Mac users: The Dockerfile is configured for cross-platform building to ensure compatibility with Heroku's x86_64 infrastructure using the --platform=linux/amd64 flag.

### API Documentation

API documentation is available through Swagger. Generate the documentation with:

```bash
make swagger
```

Then access the Swagger UI at /swagger/index.html when the application is running.

### Available Commands

Run `make help` to see all available commands:
- `make build` - Build the application
- `make run` - Run the application locally
- `make test` - Run tests
- `make test-coverage` - Run tests with coverage repor
- `make clean` - Clean up build artifacts
- `make docker-build` - Build the Docker image
- `make docker-run` - Run application in the Docker container
- `make swagger` - Generate Swagger documentation
- `make heroku-deploy` - Deploy to Heroku

### Testing Heroku Application

Application is deployed at address: https://order-packs-calculator-2025-05-39eaddfd911d.herokuapp.com/

### API Usage Examples

- **Add a new pack size**: 
```bash
curl -X POST http://localhost:8080/api/packs \
  -H "Content-Type: application/json" \
  -d '{"pack":{"size":5}}'
```

- **Remove a pack size**: 
```bash
curl -X DELETE http://localhost:8080/api/packs/5
```

- **Calculate packs for an order size**: 
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"order_size": 23}'
```

