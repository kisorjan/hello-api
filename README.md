# Simple HTTP API
This is a small Go service that returns a greeting or error depending on the first letter of the provided name.

# Endpoint
GET /hello-world?name=<string>

# Behavior
- If name starts with A–M → returns 200 OK and "Hello <name>"
- If name starts with N–Z → returns 400 Bad Request with "Invalid Input"
- If name is missing or empty → returns 400 Bad Request with "Invalid Input"

# How to Run
go run main.go

Server will start on port 9090 by default:
http://localhost:9090/hello-world?name=Alice

You can also specify a custom port:
PORT=7070 go run main.go

# Run Tests
go test -v