# Testing Notes

## Scope

This document summarizes the testing concepts and tools learned so far in the project.

## What We Tested

- health handler
- ready handler
- shorten handler
- code handler not-found path
- router-level `GET /health`
- router-level `POST /shorten`
- router-level `GET /{code}` redirect
- router-level `GET /{code}` not found
- service layer
- repository layer
- PostgreSQL repository with a real database

## Core Packages Used

### `testing`

The standard Go package for writing tests.

Main idea:
- test functions are named like `TestSomething(t *testing.T)`
- `t` is used to report failures

### `net/http/httptest`

Used for handler tests and router-level HTTP tests.

Main helpers:
- `httptest.NewRequest(...)`
- `httptest.NewRecorder()`

Main idea:
- create a fake request
- create a fake response writer
- call the handler directly
- inspect the result in memory

It is also useful for router-level tests:
- build a router
- send the request through `router.ServeHTTP(...)`
- verify route matching and HTTP behavior

### `github.com/stretchr/testify/assert`

Used to make assertions more readable.

Examples:
- `assert.Equal(...)`
- `assert.True(...)`
- `assert.False(...)`
- `assert.NotEmpty(...)`

Main idea:
- standard `testing` comes first
- `testify` is introduced only after understanding the problem it solves

## Important Concepts Learned

### Table-Driven Tests

Table-driven tests are useful when:
- many cases follow the same pattern
- only inputs and expected outputs change

They were used for simple handler method/status checks.

### Fake Dependencies

Handler tests used a fake service.
Service tests used a fake repository.

Main idea:
- test one layer at a time
- replace downstream dependencies with controlled fakes

### Direct Unit Testing

Repository tests used the real in-memory repository directly because the repository itself was the thing under test.

### Router-Level HTTP Testing

Router-level tests are one step closer to real application behavior than direct handler tests.

Main idea:
- create a Chi router
- register real paths like `/health`, `/shorten`, and `/{code}`
- send requests through `ServeHTTP(...)`
- verify the HTTP response

This tests:
- route matching
- path handling
- handler invocation through the router

### Real Database Integration Testing

The PostgreSQL repository now has integration tests that talk to a real PostgreSQL database.

Main idea:
- connect using `DATABASE_URL`
- clear the `urls` table before the test
- run real `INSERT` and `SELECT` queries through the repository
- verify the actual stored values

This is different from fake or in-memory testing because it checks:
- SQL correctness
- scanning values from PostgreSQL
- timestamp fields
- nullable fields like `expires_at`

### `t.Skip(...)`

Used when a required external dependency is not configured.

In this project:
- PostgreSQL integration tests skip if `DATABASE_URL` is not set

### `t.Cleanup(...)`

Used to register cleanup logic that runs after the test finishes.

In this project:
- close the PostgreSQL connection pool
- clear the `urls` table after each integration test

### `expected` vs `actual`

When using assertions, argument order matters.
For example:

```go
assert.Equal(t, expected, actual)
```

This makes failure messages much easier to interpret.

## Key Lesson

The project now has multiple styles of backend tests:
- handler tests
- router-level HTTP tests
- service tests
- in-memory repository tests
- PostgreSQL repository integration tests

That is important because each layer has different responsibilities and should be tested differently.
