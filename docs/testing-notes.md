# Testing Notes

## Scope

This document summarizes the testing concepts and tools learned so far in the project.

## What We Tested

- health handler
- ready handler
- shorten handler
- code handler not-found path
- service layer
- repository layer

## Core Packages Used

### `testing`

The standard Go package for writing tests.

Main idea:
- test functions are named like `TestSomething(t *testing.T)`
- `t` is used to report failures

### `net/http/httptest`

Used for handler tests.

Main helpers:
- `httptest.NewRequest(...)`
- `httptest.NewRecorder()`

Main idea:
- create a fake request
- create a fake response writer
- call the handler directly
- inspect the result in memory

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
- service tests
- repository tests

That is important because each layer has different responsibilities and should be tested differently.
