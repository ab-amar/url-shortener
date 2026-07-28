# Manual PostgreSQL Testing

This document covers the manual end-to-end testing flow for the URL shortener after wiring the app to PostgreSQL.

Before following this document, complete the local PostgreSQL setup in [postgres-local-setup.md](/Users/amarbehera/go/url-shortener/docs/postgres-local-setup.md).

## Purpose

Use this flow to verify that all of these parts work together:

- HTTP handler
- service layer
- PostgreSQL repository
- database schema

This is a manual integration test.

## Prerequisites

Before running these steps, confirm all of the following:

- PostgreSQL is installed
- PostgreSQL is running
- the `url_shortener` database exists
- the initial migration has been applied
- `DATABASE_URL` is set correctly

## 1. Start the server

Run:

```bash
export DATABASE_URL="postgres://localhost:5432/url_shortener?sslmode=disable"
go run ./cmd/server
```

What to check:

- the server starts without crashing
- you see the startup log
- port `8080` is being used unless you set `PORT`

## 2. Verify health endpoint

In a second terminal, run:

```bash
curl -i http://localhost:8080/health
```

What to check:

- status code is `200`
- response body is `Ok!`

## 3. Send a shorten request

Run:

```bash
curl -i -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/very/long/path"}'
```

What to check:

- status code is `200`
- `Content-Type` includes `application/json`
- response body contains the original URL
- response body contains a `short_code`

Example of what matters in the response:

```json
{
  "message": "...",
  "url_model": {
    "original_url": "https://example.com/very/long/path",
    "short_code": "abc12345"
  }
}
```

You do not need the exact same code value. You only need to confirm that a short code is returned.

## 4. Copy the short code

From the JSON response, note the value of `short_code`.

Example:

- if `short_code` is `abc12345`
- then your redirect test URL will be `http://localhost:8080/abc12345`

## 5. Verify the row in PostgreSQL

Run:

```bash
psql url_shortener -c "SELECT short_code, original_url, created_at FROM urls;"
```

What to check:

- a row exists for the request you just created
- `short_code` matches the API response
- `original_url` matches the URL you sent

If you want to check one code only, run:

```bash
psql url_shortener -c "SELECT short_code, original_url, created_at FROM urls WHERE short_code = 'abc12345';"
```

Replace `abc12345` with your actual code.

## 6. Test the redirect endpoint

Run:

```bash
curl -i http://localhost:8080/abc12345
```

Replace `abc12345` with your actual code.

What to check:

- status code is `302`
- `Location` header is the original URL

Example of what matters:

```text
HTTP/1.1 302 Found
Location: https://example.com/very/long/path
```

## 7. Follow the redirect automatically

Run:

```bash
curl -i -L http://localhost:8080/abc12345
```

Replace `abc12345` with your actual code.

What `-L` does:

- it tells `curl` to follow redirect responses automatically

What to check:

- `curl` follows the `302`
- final response comes from the destination URL

## 8. Test an unknown short code

Run:

```bash
curl -i http://localhost:8080/doesnotexist
```

What to check:

- status code is `404`
- response body is JSON
- error message says not found

## 9. Optional database inspection commands

Show tables:

```bash
psql url_shortener -c "\dt"
```

Describe the `urls` table:

```bash
psql url_shortener -c "\d urls"
```

Show all rows:

```bash
psql url_shortener -c "SELECT * FROM urls;"
```

## Common issues

### Server fails at startup

Possible reasons:

- `DATABASE_URL` is missing
- PostgreSQL is not running
- database name is wrong
- migration was not applied

### `POST /shorten` returns an error

Possible reasons:

- request JSON is invalid
- `Content-Type` header is missing or wrong
- database insert failed

### Redirect returns `404`

Possible reasons:

- the short code was not inserted
- the wrong short code was used
- repository lookup did not find the row

## Completion checklist

Task 65 is complete when all of these are true:

- server starts with PostgreSQL configured
- `POST /shorten` succeeds
- a row is inserted into `urls`
- `GET /{code}` returns `302`
- `Location` header matches the original URL
- unknown code returns `404`
