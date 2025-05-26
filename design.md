# Order Packs Calculator

## Overview 

The Order Packs Calculator is a service for calculating number of packs needed to ship an order.

## Definitions:

- **Pack size** is the number of items in a pack
- **Order size** is the total number of items in an order


## Principles

- The service allows users to input different pack sizes
- The service calculate the number of packs needed based on the order size and pack sizes
- Flow:
  - User inputs pack sizes
  - User removes pack sizes
  - User inputs order size
  - User presses "Calculate" button to run the calculation
  - The service calculates number of packs needed and returns the result in JSON format, and displays it on the web interface
- The service must be implemented in Golang 
- The service must be implemented as REST API service
- The service must be covered with unit tests
- The service must be implemented via 3 layer architecture: 
  - Controller: Handles HTTP requests and responses
  - Service: Contains business logic for calculating packs
  - Repository: Manages data storage (in this case, pack sizes)
- The service must be implemented using the Gin framework for HTTP handling
- A default repository implementation must use in-memory storage for pack sizes
- The service must have a simple web interface to input pack sizes and order size and a button to run the calculation of number of packs needed
- The service must have Docker support for easy deployment
- The service must have script for deploying to Heroku
- The service must public its API documentation using Swagger
- The service must have a `README.md` file with instructions on how to run the service locally, how to run tests, and how to deploy the service
- The service must have a `Makefile` with commands for running the service, running tests, building the Docker image, generating Swagger documentation, and deploying to Heroku

