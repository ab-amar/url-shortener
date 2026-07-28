# Project Architecture

This document describes the current architecture of the URL shortener service.

## Goal

Build a small backend service that is simple enough to learn from, but structured enough to discuss in an SDE2 interview.

Current concerns already covered:

- HTTP server setup
- routing
- handlers
- business logic
- persistence abstraction
- PostgreSQL integration
- middleware
- graceful shutdown
- automated tests

## High-Level Request Flow

For `POST /shorten`:

1. Chi router matches `/shorten`
2. middleware chain runs
3. handler validates and decodes the request
4. service generates the short code and timestamps
5. repository stores the URL in PostgreSQL
6. handler returns the JSON response

For `GET /{code}`:

1. Chi router extracts the path parameter
2. middleware chain runs
3. handler asks the service to resolve the code
4. service asks the repository for the stored URL
5. service rejects expired URLs
6. handler returns either:
   - `302 Found` with `Location`
   - or `404 Not Found`

## Layers

### `cmd/server`

Responsibility:
- application entry point
- configuration loading
- PostgreSQL pool creation
- dependency wiring
- HTTP server creation
- graceful shutdown

Why this layer exists:
- keeps bootstrapping concerns out of business logic

### `internal/handler`

Responsibility:
- HTTP-specific logic
- request decoding
- input validation
- writing responses
- translating service results into HTTP behavior

Examples:
- decode JSON for `POST /shorten`
- return `400`, `404`, `405`, `500`, or `302`

This layer should know:
- HTTP methods
- headers
- status codes

This layer should not know:
- SQL details
- database driver behavior

### `internal/service`

Responsibility:
- core application logic
- short code generation
- timestamp assignment
- expiry decision during resolution

Why this layer matters:
- it keeps business logic separate from transport and storage
- it is easier to test in isolation using fake repositories

### `internal/repository`

Responsibility:
- persistence access
- hide storage details behind an interface

Current implementations:

- `InMemoryRepository`
- `PostgresRepository`

Why this abstraction matters:
- handler and service do not need to know SQL
- easier unit testing with in-memory and fake dependencies
- easier migration between storage implementations

### `internal/model`

Responsibility:
- shared data structures used across layers

Current core model:
- `model.URL`

## Middleware

Current middleware stack:

1. recovery middleware
2. app header middleware
3. request ID middleware
4. logging middleware

Responsibilities:

- recover from panics and return `500`
- attach `X-App-Name`
- attach and propagate request IDs
- log request method, path, status, and request ID

## Persistence Model

Current table:
- `urls`

Current important columns:

- `id`
- `short_code`
- `original_url`
- `created_at`
- `updated_at`
- `expires_at`

Current constraints:

- primary key on `id`
- unique constraint on `short_code`

## Testing Strategy

The project now has multiple testing layers.

### Unit Tests

Used for:

- handler behavior in isolation
- service behavior with fake repositories
- in-memory repository behavior

### Router-Level HTTP Tests

Used for:

- verifying path matching and handler behavior through Chi
- checking `GET /health`
- checking `POST /shorten`
- checking redirect and not-found flows

### PostgreSQL Repository Integration Tests

Used for:

- real inserts
- real selects
- real timestamp scanning
- nullable expiry handling

These tests use:

- a real PostgreSQL database
- `DATABASE_URL`
- table cleanup before and after each test

## Design Choices

### Why use Chi now?

We started with `net/http` first to understand handlers and routing basics.

Chi was introduced after that because:

- dynamic path routing becomes awkward with the standard library alone
- path parameters like `/{code}` are cleaner with a router
- middleware composition is cleaner

### Why use a repository interface?

Because business logic should not depend directly on SQL or `pgx`.

This gives:

- testability
- cleaner separation of concerns
- flexibility to swap implementations later

### Why keep expiration logic in the service?

Expiration is a business rule, not a transport rule.

That means:

- handlers should not decide whether a link is expired
- repositories should fetch stored data
- services should interpret domain rules

## Current Limitations

The service is still intentionally simple.

Current limitations include:

- deterministic short code generation can collide
- no custom alias support
- no analytics storage yet
- no Redis caching yet
- no Docker setup yet
- no metrics endpoint yet
- no background jobs yet

These are planned future improvements, not accidental omissions.
