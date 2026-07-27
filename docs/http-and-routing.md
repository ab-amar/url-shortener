# HTTP And Routing Notes

## Scope

This document summarizes what has been learned so far about HTTP handling, routing, request flow, and API behavior in the project.

## What We Built

- a Go HTTP server using `net/http`
- handlers for `/`, `/health`, `/ready`, `/shorten`, and `/{code}`
- method checks for handlers
- JSON request decoding and JSON response writing
- redirect handling for short codes

## Core Packages Used

### `net/http`

Used for:
- `http.Server`
- `http.ResponseWriter`
- `*http.Request`
- status codes like `http.StatusOK`
- helpers like `http.Error` and `http.Redirect`

Main idea:
- handlers receive a request and build a response

### `encoding/json`

Used for:
- decoding request bodies with `json.NewDecoder(req.Body).Decode(...)`
- encoding JSON responses with `json.NewEncoder(w).Encode(...)`

Main idea:
- Go structs map cleanly to JSON request and response bodies

### `net/url`

Used for:
- parsing and validating the input URL before shortening

Main idea:
- application code should not trust raw input strings

### `github.com/go-chi/chi/v5`

Used for:
- clean route registration
- dynamic route params like `/{code}`
- middleware registration with `router.Use(...)`

Main idea:
- the standard library teaches the basics first, then a router becomes useful once dynamic routes and middleware become awkward to manage manually

## Important Concepts Learned

### Handler

A handler is a function or type that processes an HTTP request and writes a response.

Example shape:

```go
func SomeHandler(w http.ResponseWriter, req *http.Request)
```

### Method Checking

Handlers should explicitly reject unsupported methods.
This keeps API behavior predictable and avoids accidental misuse of endpoints.

### Content Type

The response should declare whether it is:
- plain text
- JSON

This is done with:

```go
w.Header().Set("Content-Type", "...")
```

### Status Codes

Handlers should return meaningful HTTP status codes such as:
- `200 OK`
- `400 Bad Request`
- `404 Not Found`
- `405 Method Not Allowed`
- `500 Internal Server Error`

### Liveness vs Readiness

`/health` answers: is the process alive?
`/ready` answers: is the service ready to receive traffic?

Right now they are both simple, but later `/ready` can become stricter when database and cache dependencies are added.

## API Behavior Added So Far

### `POST /shorten`

- accepts JSON input
- validates URL format
- returns JSON response
- returns consistent JSON error responses for API failures

### `GET /{code}`

- resolves a short code
- redirects to the original URL if found
- returns JSON `404` if not found

## Key Lesson

The service already has the beginnings of real API design:
- explicit routes
- explicit validation
- explicit status codes
- separate operational endpoints
- predictable response formats
