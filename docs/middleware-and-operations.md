# Middleware And Operations Notes

## Scope

This document summarizes what has been learned so far about middleware, logging, request lifecycle operations, and server hardening.

## Middleware Concepts Learned

Middleware is code that wraps a handler.

Instead of putting common logic inside every handler, middleware applies that logic around all requests.

Conceptual shape:

```go
func SomeMiddleware(next http.Handler) http.Handler
```

Main idea:
- request enters middleware
- middleware calls `next`
- middleware can do work before and after the handler

## Middleware Added So Far

### `AppHeaderMiddleware`

Adds a response header:
- `X-App-Name: url-shortener`

This was used to make middleware chaining easy to observe.

### `RequestIDMiddleware`

Generates a request ID, stores it in:
- response header `X-Request-ID`
- request context

This makes request logs easier to correlate.

### `LoggingMiddleware`

Logs:
- method
- path
- status
- request ID

This is the first centralized request observability layer.

### `RecoveryMiddleware`

Catches panics from downstream handlers, logs the panic, and returns a controlled JSON `500` response.

Main idea:
- `recover()` only works inside a deferred function
- middleware is the right place for panic recovery

## Structured Logging

The project uses `log/slog` from the standard library for lifecycle and request logging.

Lifecycle logs added so far:
- server starting
- shutdown signal received
- server failed
- server shutdown failed
- server stopped

Main lesson:
- structured key/value logs are easier to search and reason about than plain prints

## Server Hardening Added So Far

### Graceful Shutdown

The server already supports:
- OS signal handling
- shutdown timeout with `context.WithTimeout`
- controlled stop behavior

### Server Timeouts

The server now has:
- `ReadHeaderTimeout`
- `ReadTimeout`
- `WriteTimeout`
- `IdleTimeout`

Main meanings:
- `ReadHeaderTimeout`: how long to wait for request headers
- `ReadTimeout`: how long to read the full request
- `WriteTimeout`: how long to write the response
- `IdleTimeout`: how long to keep an idle connection open

## Operational Endpoints

The project now distinguishes:
- `/health` for liveness
- `/ready` for readiness

This is small now, but it becomes more meaningful once real dependencies like PostgreSQL are introduced.

## Key Lesson

The project is no longer just about request/response code.
It now includes real service operation concerns:
- observability
- panic containment
- request correlation
- connection limits
- liveness and readiness
