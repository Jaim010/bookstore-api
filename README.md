# Golang 1.18 & Gin Bookstore API

## Setup
Before running the API, create a `.env` file and place it in `/config/`. Copy the content below into `/config/.env` and change each value to match your credentials and database. 
```
POSTGRES_USER=db-user
POSTGRES_PASSWORD=user-pwd
POSTGRES_DB=db-name
POSTGRES_URL=url
POSTGRES_PORT=5432
```

## Building
To build the API, run the command below in the `root` directory (`~/`). <br />
```go build -o bin/main ./cmd/main.go```
or <br />
```make build```

## Running
To run the API, run the command below in the `cmd` directory (`/cmd/`). <br />
```go run .```
or run `make run` in the `root` directory (`~/`). <br />


## Testing
To run the tests, run the command below in the `tests` directory (`/tests/`) <br />
```go test -v``` <br /> 
Or run `go test -v ./...` in the root directory (`/`).
Or run `make test` in the root directory (`/`).
