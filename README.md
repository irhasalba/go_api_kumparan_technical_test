# Skill Test Backend Engineer Kumparan

## Requirement
1. Go >=1.23
2. install makefile
3. Docker

## Background
This project was created as part of the recruitment process at Kumparan. The project is quite simple: building an API to add an article and then display it.

## Solution
* I built an API using Golang Fiber and SQLC as the SQL compiler, and I used PostgreSQL as the database.
* I used a simple approach with limit and offset to solve problem number 1. I admit this is not a long-term solution, but I chose it due to time and resource constraints.
* For point number 2, I chose to use a rate limiter to handle excessive user requests at the same time. This method is simple but effective for completing the task in a short time.

## How to Run the Project
1. After cloning this repository, run the following command:
   `go mod tidy`
2. Run docker compose to set up the project environment:
   `docker compose up`
3. run command this to copy env :
`make env`
4. Then run:
   `make migrate-up` to apply the database migrations
5. Next, run:
   `go run main.go`
6. Finally, access `http://localhost:3031/`. If you see "hello, world", it means the application is running and

## Running Integration Test
For run integration test just copy command below :
`make integration-test`