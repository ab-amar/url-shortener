# Go Backend Fundamentals

This document summarizes the concepts, packages, and patterns used in the first category of the URL shortener project.

## Scope of This Category

This category covered the basics of building and running a Go HTTP service using the standard library.

Completed topics:

1. Starting an HTTP server with `net/http`
2. Separating handlers from `main.go`
3. Basic routing with `http.ServeMux`
4. Writing HTTP responses
5. Handling HTTP methods
6. Basic configuration
7. Reading configuration from environment variables
8. Validating configuration at startup
9. Graceful shutdown
10. Request lifecycle basics
11. `context.Context` basics

## High-Level Architecture So Far

Current structure:

- `cmd/server/main.go`
  - application entrypoint
  - config loading
  - server creation
  - startup
  - graceful shutdown

- `internal/handler/handler.go`
  - HTTP handlers
  - request checks
  - response writing

- `internal/config/config.go`
  - config struct
  - env loading
  - startup validation

## Core Go Concepts Learned

### 1. `package main`

`package main` is used for executable programs in Go.

Why it matters:

- this is where the process starts
- `main()` is the entrypoint
- backend services usually keep `main()` small and focused on wiring

### 2. Functions

Go programs are built from functions.

Examples used so far:

- `main()`
- `createServer(port string) http.Server`
- handler functions like:
  - `HealthHandler(w http.ResponseWriter, req *http.Request)`

Why this matters:

- functions define behavior
- helper functions reduce complexity in `main()`

### 3. Structs

A struct groups related data together.

Example:

```go
type Config struct {
	Port string
}
```

Why it matters:

- configuration belongs together
- structs scale better than scattered constants

### 4. Exported vs unexported names

In Go:

- names starting with uppercase letters are exported
- names starting with lowercase letters are package-private

Examples:

- exported:
  - `Config`
  - `Port`
  - `NewConfig`
  - `HealthHandler`

- unexported:
  - `createServer`

Why it matters:

- package boundaries are enforced through naming
- `main` can only access exported names from other packages

### 5. Errors

Go uses returned `error` values instead of exceptions for normal error handling.

Examples used:

- `server.ListenAndServe()`
- `server.Shutdown(ctx)`
- `config.NewConfig()`
- `strconv.Atoi(...)`

Common pattern:

```go
value, err := someFunction()
if err != nil {
	panic(err)
}
```

Why it matters:

- startup and shutdown failures must be handled explicitly
- invalid config should fail fast

### 6. Goroutines

A goroutine runs a function concurrently.

Example shape:

```go
go func() {
	// background work
}()
```

Why it mattered here:

- `ListenAndServe()` blocks
- graceful shutdown required the server to run in the background while `main()` waited for a signal

### 7. Channels

A channel is used to send values between parts of a program.

Example shape:

```go
sigChan := make(chan os.Signal, 1)
```

Important operations:

- send:
  - `ch <- value`
- receive:
  - `<-ch`

Why it mattered here:

- `os/signal` sends shutdown signals into a channel
- `main()` waits on that channel before starting graceful shutdown

### 8. Context

`context.Context` carries control information for a unit of work.

Main uses so far:

- shutdown timeout:
  - `context.WithTimeout(context.Background(), ...)`
- request lifecycle:
  - `req.Context()`

Why it matters:

- work can be canceled
- work can have deadlines
- request-scoped work should respect request lifetime

## HTTP Concepts Learned

### 1. HTTP server

An HTTP server:

- listens on a port
- accepts requests
- routes them to handlers
- writes responses

In Go, `http.Server` is the core server type.

### 2. Handler

A handler is code that processes an HTTP request and writes an HTTP response.

Current handlers:

- `RootHandler`
- `HealthHandler`

Handler signature:

```go
func SomeHandler(w http.ResponseWriter, req *http.Request)
```

### 3. `http.ResponseWriter`

`http.ResponseWriter` is used to build the HTTP response.

Used for:

- setting headers
- setting status code
- writing response body

Examples:

```go
w.Header().Set("Content-Type", "text/plain")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "Ok!")
```

### 4. `*http.Request`

`*http.Request` contains incoming request data.

Used so far for:

- checking method:
  - `req.Method`
- accessing request context:
  - `req.Context()`

### 5. Routing with `http.ServeMux`

`http.ServeMux` maps paths to handlers.

Used routes:

- `/`
- `/health`

Important routing lesson:

- `/` is broad
- `/health` is more specific
- more specific routes win

### 6. Method handling

Handlers should explicitly check allowed methods.

Used pattern:

```go
if req.Method != http.MethodGet {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
```

Why it matters:

- route existence and method validity are different concerns
- `405 Method Not Allowed` is different from `404 Not Found`

### 7. Response structure

An HTTP response has three important parts:

1. status code
2. headers
3. body

Example:

```go
w.Header().Set("Content-Type", "text/plain")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "Ok!")
```

### 8. Common status codes used so far

- `200 OK`
  - successful request
- `405 Method Not Allowed`
  - route exists but method is not allowed

### 9. Headers used so far

#### `Content-Type`

Tells the client what kind of response body is being returned.

Used value:

- `text/plain`

### 10. `http.Error`

`http.Error(...)` is a helper that writes an error response.

Used for:

- rejecting unsupported methods

Why it matters:

- fast way to send a status code plus error body

## Configuration Concepts Learned

### 1. What configuration is

Configuration is runtime input that changes how a service runs.

Used example:

- server port

Why it is different from business logic:

- business logic defines service behavior
- configuration defines runtime setup

### 2. Config package

Config loading was placed in:

- `internal/config/config.go`

Why:

- keeps config sourcing and validation out of `main.go`
- gives one place for runtime settings

### 3. Environment variables

Used:

- `PORT`

Loading pattern:

```go
port := os.Getenv("PORT")
if port == "" {
	port = "8080"
}
```

### 4. Validation

Config values should be validated at startup.

Used validation:

- parse port using `strconv.Atoi(...)`

Why:

- invalid runtime config should fail before the app starts serving requests

## Graceful Shutdown Concepts Learned

### 1. Why graceful shutdown matters

Backend services should not stop abruptly if they can avoid it.

Graceful shutdown means:

- stop accepting new requests
- let current work finish
- exit cleanly within a timeout

### 2. Signal handling

Used packages:

- `os`
- `os/signal`

Used signal:

- `os.Interrupt`

Why:

- lets the app react to `Ctrl + C`

### 3. Shutdown context

Used:

```go
ctx, cancel := context.WithTimeout(context.Background(), ...)
defer cancel()
```

Why:

- shutdown should not wait forever

### 4. `http.ErrServerClosed`

During normal graceful shutdown, `ListenAndServe()` returns `http.ErrServerClosed`.

Why it matters:

- this is expected
- it should not be treated like a real crash

## Request Lifecycle Concepts Learned

### 1. Request lifecycle

For a request like `/health`:

1. client sends request
2. server accepts it
3. mux matches route
4. handler runs
5. response is written
6. request ends

### 2. Request context

Every request has a context:

```go
ctx := req.Context()
```

Why it matters:

- later this context will be passed to service and repository layers
- if the request is canceled, deeper work can stop too

### 3. Difference between request context and background context

`context.Background()`

- root context
- used at top-level app code like shutdown

`req.Context()`

- tied to one incoming HTTP request
- used during request handling

## Packages Used So Far

### `net/http`

Used for:

- `http.Server`
- `http.NewServeMux()`
- `http.HandleFunc` via mux
- `http.ResponseWriter`
- `*http.Request`
- `http.MethodGet`
- `http.StatusOK`
- `http.StatusMethodNotAllowed`
- `http.Error`
- `http.ErrServerClosed`
- `server.ListenAndServe()`
- `server.Shutdown(...)`

### `fmt`

Used for:

- writing plain-text response bodies with `fmt.Fprintf(...)`

### `os`

Used for:

- reading environment variables with `os.Getenv(...)`
- interrupt signal type through `os.Interrupt`

### `os/signal`

Used for:

- receiving OS interrupt signals using `signal.Notify(...)`

### `context`

Used for:

- `context.Background()`
- `context.WithTimeout(...)`

### `time`

Used for:

- shutdown timeout duration

### `strconv`

Used for:

- validating numeric port input with `strconv.Atoi(...)`

## Design Principles Learned So Far

### 1. Keep `main()` small

`main()` should mostly:

- load config
- construct top-level objects
- start the app
- handle shutdown

### 2. Keep handlers transport-focused

Handlers should deal with:

- HTTP method checks
- request/response behavior
- request context

Not with:

- business logic
- storage logic
- unrelated startup code

### 3. Introduce complexity gradually

So far the project intentionally avoided:

- third-party routers
- databases
- Redis
- middleware chains
- service/repository abstractions beyond what was needed

Why:

- fundamentals should be understood before abstractions are added

### 4. Fail fast on invalid config

Bad config should be caught during startup, not after the app starts serving traffic.

### 5. Use the standard library first

So far everything was built with the Go standard library.

Why:

- learn the primitives first
- understand the problems before adding tools

## What This Category Prepared Us For

This category prepared the project for the next one:

- adding `POST /shorten`
- reading JSON requests
- defining request/response structs
- validating input
- introducing model, service, and repository boundaries

The main outcome is that the project now has a working, structured, configurable, and gracefully shutting down HTTP service built with standard Go tools.
