# DigitalExchange

A modern digital exchange system built with Go, utilizing Wire for dependency injection and Cobra for command-line interface management. This project provides a RESTful API for managing exchange operations, integrated with a PostgreSQL database.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Development](#development)
- [Running with Docker](#running-with-docker)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Overview
DigitalExchange is a backend application designed to handle exchange-related operations such as order management, balance tracking, and order book updates. It leverages:
- **Go** as the primary programming language.
- **Wire** for dependency injection to manage service and repository dependencies efficiently.
- **Cobra** to provide a command-line interface for running and managing the application.
- **PostgreSQL** as the database backend, managed via Docker.

The project structure includes API controllers, models, database migrations, and services, making it extensible and maintainable.

## Features
- RESTful API for creating and managing orders.
- Balance management for different assets.
- Order book tracking for exchange markets.
- CLI commands for application control using Cobra.
- Dockerized deployment with PostgreSQL integration.

## Prerequisites
- **Go** (version 1.20 or higher)
- **Docker** and **Docker Compose** (for running the application and database)
- **Git** (for cloning the repository)
- PostgreSQL client (optional, for manual database management)

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/DigitalExchange.git
   cd DigitalExchange
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Install Wire (for dependency injection):
   ```bash
   go install github.com/google/wire/cmd/wire@latest
   ```
4. Install Cobra (for CLI):
   ```bash
   go install github.com/spf13/cobra-cli@latest
   ```

## Usage
### Running the Application
1. Ensure the `.env` file is configured with the correct environment variables (see [Configuration](#configuration)).
2. Run the application using the Cobra CLI:
   ```bash
   go run main.go server
   ```
   This starts the API server on the port specified in `.env` (default: 8443).

### Available CLI Commands
Use `go run main.go --help` to see available commands. Example:
- `go run main.go server`: Starts the API server.
- `go run main.go migrate`: Runs database migrations.

## Configuration
Copy the `.env.example` file to `.env` and update the following variables:
```env
# App
APP_NAME=exchange
APP_ENV=local
APP_PORT=8443
APP_HOST=localhost
APP_TZ=Asia/Tehran

# Database
DB_DRIVER=postgres
DB_HOST=postgres
DB_PORT=5431
DB_DATABASE="DigitalExchange-api"
DB_USERNAME="DigitalExchange"
DB_PASSWORD="myawesomepassword"
DB_SSL_MODE=disable
DB_MAX_OPEN_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=50
DB_CONNECTION_MAX_LIFETIME=3600
```
- **Note:** Set `DB_HOST=postgres` to connect to the Dockerized PostgreSQL service.

## Development
### Running Locally
1. Start the development environment:
   ```bash
   go run main.go server
   ```
2. Test the API using tools like Postman or `curl`.

### Generating Wire Injections
Run the following command to generate Wire bindings:
```bash
wire
```
Ensure the `wire.go` file is updated in the appropriate package (e.g., `wire_gen.go`).

### Database Migrations
Run migrations to set up the database schema:
```bash
go run main.go migrate
```

## Running with Docker
1. Build and start the services:
   ```bash
   docker-compose up --build
   ```
2. Access the API at `http://localhost:8443` and the database at `localhost:5431`.

### Stopping the Services
```bash
docker-compose down
```

## API Documentation
The API includes endpoints for order management. Example request (using Postman):
- **Endpoint:** `POST /api/exchange/order/create`
- **Body:**
  ```json
  {
    "exchange_name": 1,
    "client_order_id": "buy-btc-limit-20250530-001",
    "symbol": "BTCUSDT",
    "side": "buy",
    "type": "limit",
    "quantity": 0.2,
    "price": 60000.0,
    "time_in_force": "GTC"
  }
  ```
- Refer to the `api/controllers` package for more endpoints.

## Contributing
1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch`.
3. Make your changes and commit: `git commit -m "Add new feature"`.
4. Push to the branch: `git push origin feature-branch`.
5. Submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

### نکات
- این فایل رو با نام `README.md` توی ریشه پروژه (کنار `docker-compose.yml`) ذخیره کن.
- لینک GitHub (`https://github.com/yourusername/DigitalExchange.git`) رو با آدرس واقعی مخزن‌ات جایگزین کن.
- اگه `.env.example` داری، مطمئن شو که توی پروژه هست و توی `README` بهش اشاره شده.
- می‌تونی بخش‌های اضافی مثل "Testing" یا "Troubleshooting" رو اضافه کنی اگه نیاز داری.

اگه می‌خوای تغییر خاصی بدی (مثلاً اضافه کردن دستورات بیشتر یا تغییر سبک)، بگو تا ویرایش کنم! 😊