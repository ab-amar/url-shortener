# URL Shortener

A production-oriented URL shortener backend in Go, built incrementally to learn backend engineering, Go fundamentals, HTTP services, persistence, testing, and system design.

## Current State

The project currently supports:

- `POST /shorten` to create a short code for a URL
- `GET /{code}` to redirect to the original URL
- `GET /health` for health checks
- `GET /ready` for readiness checks
- PostgreSQL-backed persistence

## Prerequisites

- Go installed
- PostgreSQL installed and running

If PostgreSQL is not set up yet, follow [docs/postgres-local-setup.md](/Users/amarbehera/go/url-shortener/docs/postgres-local-setup.md).

## Quick Start

From the project root:

```bash
go test ./...
```

Set the required environment variable:

```bash
export DATABASE_URL="postgres://localhost:5432/url_shortener?sslmode=disable"
```

Optionally set the port:

```bash
export PORT=8080
```

Run the server:

```bash
go run ./cmd/server
```

## Manual API Test

In another terminal:

Check health:

```bash
curl -i http://localhost:8080/health
```

Create a short URL:

```bash
curl -i -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'
```

Test redirect:

```bash
curl -i http://localhost:8080/<short-code>
```

For the full manual testing flow, use [docs/manual-postgres-testing.md](/Users/amarbehera/go/url-shortener/docs/manual-postgres-testing.md).

## Database Migrations

Apply migrations manually:

```bash
psql url_shortener -f migrations/000001_create_urls_table.up.sql
psql url_shortener -f migrations/000002_add_url_metadata_columns.up.sql
```

If you already created the table earlier, you only need to apply:

```bash
psql url_shortener -f migrations/000002_add_url_metadata_columns.up.sql
```

## Project Docs

- [docs/postgres-local-setup.md](/Users/amarbehera/go/url-shortener/docs/postgres-local-setup.md)
- [docs/manual-postgres-testing.md](/Users/amarbehera/go/url-shortener/docs/manual-postgres-testing.md)
- [docs/fundamentals.md](/Users/amarbehera/go/url-shortener/docs/fundamentals.md)
- [docs/http-and-routing.md](/Users/amarbehera/go/url-shortener/docs/http-and-routing.md)
- [docs/testing-notes.md](/Users/amarbehera/go/url-shortener/docs/testing-notes.md)
