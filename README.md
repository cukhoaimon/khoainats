# KhoaiNATS

KhoaiNATS is a side project that demonstrates a modern, modular backend architecture using Go, NATS, and Gin. It is designed for scalability, maintainability, and ease of development, leveraging best practices in Go backend development.

## Features

- **Modular Project Structure**: Clear separation of concerns with `internal/api`, `internal/repositories`, and `internal/services`.
- **OpenAPI-Driven Development**: API definitions and server stubs are generated from OpenAPI specifications, ensuring consistency and rapid iteration.
- **Gin Web Framework**: Fast, minimalist HTTP server using [Gin](https://github.com/gin-gonic/gin).
- **NATS Integration**: Uses [NATS](https://nats.io/) for high-performance messaging and event-driven architecture. Mainly use it for MQTT, I intentionally develop this repo as an IoT platform.
- Simple SES Server

## Technologies Used

- **Go**: Main programming language.
- **Gin**: HTTP web framework.
- **NATS**: Embedded as a third-party dependency for messaging.
- **PostgreSQL/PostGIS**: Relational database with spatial extensions.
- **OpenAPI**: API contract and code generation.

## Getting Started

1. **Clone the repository**

   ```bash
   git clone https://github.com/cukhoaimon/khoainats.git
   cd khoainats
   ```

2. **Start dependencies with Docker Compose**

   ```bash
   docker-compose up -d
   ```

3. **Run the server**
   ```bash
   go run cmd/api/main.go
   ```

## Project Structure

### `/api`

OpenAPI specs with generated server stub and client code in `__generated__`

### `/cmd`

Entry-points for this project, it contains NATS server, service platform,...

### `/internal`

Private application and library, I don't use `/pkg` until I found something large enough that need to expose into miniservice, microservice or nanoservice.

### `/third-party`

External helper tools, forked code and other 3rd party utilities (e.g., NATS).

### `/tools`

Supporting tools for this project. These tools can import code from the /pkg and /internal directories.
