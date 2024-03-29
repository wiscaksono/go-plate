# My Go Project with Fiber and GORM

This is a template for organizing a Go project using the Fiber web framework and GORM ORM.

## Project Structure

```plaintext
go/
|-- cmd/
|   |-- myapp/
|       |-- main.go
|
|-- internal/
|   |-- app/
|       |-- handler/
|       |   |-- handler1.go
|       |   |-- handler2.go
|       |
|       |-- middleware/
|       |   |-- middleware1.go
|       |   |-- middleware2.go
|       |
|       |-- model/
|       |   |-- model1.go
|       |   |-- model2.go
|       |
|       |-- repository/
|       |   |-- repository1.go
|       |   |-- repository2.go
|       |
|       |-- router/
|           |-- router.go
|
|-- config/
|   |-- config.go
|
|-- go.mod
|-- go.sum
|-- README.md
```

## Overview

- **cmd**: Contains the main application. main.go initializes and runs the application.
- **internal/app**: Main application logic.
  - **handler**: Handles HTTP requests and responses.
  - **middleware**: Custom middleware functions.
  - **model**: Data models.
  - **repository**: Database access logic using GORM.
  - **router**: Initializes and configures the router.
- **scripts**: Project-related scripts (e.g., setup and seeding scripts).
- **config**: Configuration settings for the application.

## Usage

1. Clone the repository:

```bash
git clone https://github.com/wiscaksono/go-plate.git
```

2. Install dependencies:

```bash
cd myproject
go mod download
```

3. Run the application

```bash
go run cmd/myapp/main.go
```

## Database Migrations

Apply database migrations using GORM:

```bash
go run cmd/myapp/main.go migrate
```

## Contributing

Feel free to contribute by opening issues or pull requests.
