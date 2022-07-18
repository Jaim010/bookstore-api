# Golang-Gin Bookstore API

## Setup
Before running the API, create a `.env` file and place it in `/config/`. Copy the content below into `/config/.env` and change each value to match your credentials and database. 
```
POSTGRES_USER=db-user
POSTGRES_PASSWORD=user-pwd
POSTGRES_DB=db-name
POSTGRES_URL=url
POSTGRES_PORT=5432
```

## Running the API
To run the API, enter the command below in the `cmd` directory (`/cmd/`). <br />
```go run .```

## Running tests
To run the tests, enter the command below in the `tests` directory (`/tests/`) <br />
```go test -v``` <br /> 
Or run `go test -v ./...` in the root directory (`/`).