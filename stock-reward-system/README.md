# Stock Reward System

## Overview
The Stock Reward System is a web application built using the Gin framework in Go. It provides a RESTful API for managing stock rewards, allowing users to record rewards, fetch stock data, and retrieve user statistics.

## Features
- Record stock rewards
- Fetch today's stock prices
- Retrieve historical INR values
- Access user statistics

## Project Structure
```
stock-reward-system
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── api
│   │   ├── handler.go   # API handler functions
│   │   └── routes.go    # API route definitions
│   ├── db
│   │   ├── database.go   # Database connection and setup
│   │   └── schema.sql    # SQL schema for the database
│   ├── model
│   │   └── reward.go     # Data structures for stock rewards
│   ├── service
│   │   └── reward_service.go # Business logic for the reward system
│   └── util
│       └── error.go      # Utility functions for error handling
├── go.mod                # Go module definition
├── go.sum                # Module dependency checksums
└── README.md             # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd stock-reward-system
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up the database:
   - Create a PostgreSQL database and update the connection details in `internal/db/database.go`.
   - Run the SQL schema:
     ```
     psql -U <username> -d <database> -f internal/db/schema.sql
     ```

4. Run the application:
   ```
   go run cmd/main.go
   ```

## API Specifications
### Endpoints
- `POST /rewards` - Record a new stock reward
- `GET /stocks/today` - Fetch today's stock prices
- `GET /inr/historical` - Retrieve historical INR values
- `GET /user/stats` - Access user statistics

## Edge Case Handling
The application includes error handling for various edge cases, such as:
- Invalid input data
- Database connection errors
- Not found errors for resources

## License
This project is licensed under the MIT License.