# User Details gRPC Service

This project is a Golang gRPC service that provides functionalities for managing user details and includes a search capability.


## Overview

The service simulates a database by maintaining a list of user details within a variable. It provides gRPC endpoints for fetching user details based on a user ID, retrieving a list of user details based on a list of user IDs, and implementing a search functionality to find user details based on specific criteria.

## Features

- Fetch user details by user ID
- Retrieve a list of user details by a list of user IDs
- Search user details based on criteria (e.g., city, phone number, marital status)
- Comprehensive unit tests for all endpoints
- Dockerized for easy deployment

## Prerequisites

- Go 1.16+
- Protocol Buffers compiler (protoc)
- Docker (for containerization)

## Installation

### Clone the Repository

```sh
git clone https://github.com/yourusername/user-details-grpc-service.git
cd user-details-grpc-service
```

## Build the Application
  You can build the application using the provided build.sh script:

```sh
./build.sh
```
## Usage
### Run the Application
Start the gRPC server:

```sh
./build/userService
```
## Accessing gRPC Endpoints

You can use any gRPC client to interact with the service. Below are examples using `grpcurl` to interact with the various endpoints.

### Fetch User Details by User ID

Fetch user details by ID:

```sh
grpcurl -plaintext -d '{"id": "1"}' localhost:50051 userdetails.UserService/GetUserDetails
```
### Retrieve a List of User Details by a List of User IDs
```sh
grpcurl -plaintext -d '{
  "pageNumber": 1,
  "pageSize": 10,
  "ids": ["1", "2", "3"]
}' localhost:50051 userdetails.UserService/GetUserList
```

### Search User Details
```sh
grpcurl -plaintext -d '{
  "pageNumber": 1,
  "pageSize": 10,
  "filters": {
    "city": "LA",
    "married": "TRUE"
  }
}' localhost:50051 userdetails.UserService/SearchUsers
```

## Testing
Run the unit tests using the following command:

```sh
go test ./...
```
## Dockerization
### Build the Docker Image
```sh
docker build -t user-details-grpc-service .
```
##Run the Docker Container
```sh
docker run -p 50051:50051 user-details-grpc-service
```




