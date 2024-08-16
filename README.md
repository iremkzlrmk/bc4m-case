# bc4m-case

This project is a simple Go API using the Fiber framework, equipped with Swagger documentation and Docker support. The API includes endpoints for basic operations, health checks, and echoing request data.

## Features

- **GET /**: Returns a static message `{"msg":"BC4M"}`.
- **GET /health**: Provides the health status of the application.
- **POST /**: Echoes the data posted in the request body.
- **Swagger Documentation**: Available at `/swagger/index.html` for API documentation.
