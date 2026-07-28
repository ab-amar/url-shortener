# Observability

Observability is how we understand the behavior of a running service from the signals it produces.

For a backend service, the three common observability pillars are:

- logs
- metrics
- traces

## Logs

Logs are event records written by the application.

Examples:

- server started
- server stopped
- request completed
- panic recovered
- database connection failed

In this project, logs are currently written using `log/slog`.

Logs are useful when:

- debugging a specific request
- understanding an error message
- looking at detailed request context

## Metrics

Metrics are numeric measurements collected over time.

Examples:

- total number of requests
- total number of errors
- total number of redirects
- total number of shorten requests
- request duration

Metrics are useful when:

- identifying trends
- building dashboards
- setting alerts
- understanding service health over time

## Traces

Traces follow a single request through different parts of a system.

In a distributed system, a trace can show:

- request entered the API
- request hit the service layer
- request queried PostgreSQL
- request queried Redis
- request called another internal service

Traces are useful when:

- debugging latency
- understanding where time is spent
- following a single request across many components

This project does not have tracing yet.

## What This Project Already Has

The project already has some observability foundations:

- structured logs with `slog`
- request completion logging in middleware
- panic logging in recovery middleware
- startup and shutdown logs in `main.go`
- database connection failure logs
- request IDs for correlating logs
- `/health` endpoint
- `/ready` endpoint

These are useful, but they are still mostly log-based observability.

## What Is Missing

The project does not yet have:

- request counters
- error counters
- endpoint-specific counters
- redirect counters
- shorten counters
- request latency measurements
- database query metrics
- tracing

## Why Metrics Are The Next Step

Metrics are the next best improvement for this project because:

- the service is still small
- metrics give production value quickly
- counters are easier to add than tracing
- they make service behavior measurable over time

That is why the next task is to add basic counters for important API behavior.
