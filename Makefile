.PHONY: all build run test test-coverage clean docker-build docker-run swagger help deploy-heroku

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=packs-calculator
MAIN_PATH=./cmd/server/main.go

# Docker parameters
DOCKER_IMG=order-packs-calculator
DOCKER_TAG=latest
DOCKER_PORT=8080

# Heroku parameters
HEROKU_APP_NAME=order-packs-calculator-2025-05

all: test build

build:
	@echo "Building..."
	$(GOBUILD) -o ./bin/$(BINARY_NAME) $(MAIN_PATH)

run:
	@echo "Running..."
	$(GORUN) $(MAIN_PATH)

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -cover -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)
	rm -f coverage.out coverage.html

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMG):$(DOCKER_TAG) .

docker-run:
	@echo "Running Docker container..."
	docker run -p $(DOCKER_PORT):$(DOCKER_PORT) --name $(DOCKER_IMG) $(DOCKER_IMG):$(DOCKER_TAG)

swagger:
	@echo "Generating Swagger documentation..."
	@if ! command -v swag &> /dev/null; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	swag init -g $(MAIN_PATH) -o docs

deploy-heroku:
	@echo "Deploying to Heroku..."
	@if ! command -v heroku &> /dev/null; then \
		echo "Heroku CLI not found. Please install it first."; \
		exit 1; \
	fi
	heroku container:login
	heroku stack:set container --app order-packs-calculator-2025-05
	heroku container:push web --app $(HEROKU_APP_NAME)
	heroku container:release web --app $(HEROKU_APP_NAME)
	@echo "Deployed to: https://$(HEROKU_APP_NAME).herokuapp.com"

help:
	@echo "Available commands:"
	@echo "  make build            - Build the application binary"
	@echo "  make run              - Run the application locally"
	@echo "  make test             - Run tests"
	@echo "  make test-coverage    - Run tests with coverage report"
	@echo "  make clean            - Clean up build artifacts"
	@echo "  make docker-build     - Build Docker image"
	@echo "  make docker-run       - Run application in Docker container"
	@echo "  make swagger          - Generate Swagger documentation"
	@echo "  make deploy-heroku    - Deploy to Heroku"
	@echo "  make help             - Show this help message"

# Default target
.DEFAULT_GOAL := help