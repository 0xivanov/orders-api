
# Orders API

The Orders API is a Go-based web service that manages orders for an online store. It uses Redis as a backend for storing order data.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/0xivanov/orders-api.git
cd orders-api
```

### Running the API

1. Start the Redis container:

```bash
docker-compose up -d redis
```

2. Build and run the Orders API:

```bash
docker-compose up orders-api
```

The API will be accessible at [http://localhost:3000](http://localhost:3000).

## API Endpoints

- **GET /orders**: Retrieve a list of orders.
- **GET /orders/{orderId}**: Retrieve details for a specific order.
- **POST /orders**: Create a new order.
- **DELETE /orders/{orderId}**: Delete a specific order.

## Dependencies

The API relies on the following third-party libraries:

- [github.com/cespare/xxhash/v2](https://pkg.go.dev/github.com/cespare/xxhash/v2)
- [github.com/dgryski/go-rendezvous](https://pkg.go.dev/github.com/dgryski/go-rendezvous)
- [github.com/go-chi/chi/v5](https://pkg.go.dev/github.com/go-chi/chi/v5)
- [github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid)
- [github.com/redis/go-redis/v9](https://pkg.go.dev/github.com/redis/go-redis/v9)
