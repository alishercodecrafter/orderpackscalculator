#!/bin/bash
# scripts/generate_swagger.sh

# Ensure swag is installed
if ! command -v swag &> /dev/null; then
    echo "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

echo "Generating Swagger documentation..."
swag init -g cmd/server/main.go -o docs

echo "Swagger documentation generated in docs/ directory"