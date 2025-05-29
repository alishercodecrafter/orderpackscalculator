# Order Packs Calculator

## Overview 

The Order Packs Calculator is a service for calculating number of packs needed to ship an order.

## Definitions:

- **Pack size** is the number of items in a pack
- **Order size** is the total number of items in an order


## Principles

- There are 3 rules calculation must follow:
  - #1. Only whole packs can be sent. Packs cannot be broken open
  - #2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order
  - #3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each
     order (Please note, rule #2 takes precedence over rule #3)
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

## Implementation

### File Structure

```
order-packs-calculator/
├── cmd/
│   └── server/
│       └── main.go             # Entry point
├── docs/
│   └── docs.go                 # Swagger documentation
│   └── swagger.json            # Swagger JSON file
│   └── swagger.yaml            # Swagger YAML file
├── internal/
│   ├── controller/
│   │   └── controller.go       # HTTP handlers
│   ├── service/
│   │   └── service.go          # Business logic
│   │   └── service_test.go     # Unit tests for service
│   │   └── mock_repository.go  # Mock for repository
│   ├── repository/
│   │   └── mem_impl.go         # In-memory implementation
│   └── model/
│       └── model.go            # Data models
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   └── styles.css      # CSS styles
│   │   └── js/
│   │       └── main.js         # Frontend JavaScript
│   └── templates/
│       └── index.html          # HTML template
├── .gitignore
├── Dockerfile                  # Dockerfile for building the image
├── go.mod                      # Go module file
│   └── go.sum
├── design.md                   # Design document
├── heroku.yml                  # Heroku deployment configuration
├── Makefile                    # Makefile for build and run commands
└── README.md                   # Project documentation
```