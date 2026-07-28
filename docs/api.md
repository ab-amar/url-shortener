# API Documentation

This document describes the current HTTP API exposed by the URL shortener service.

## Base URL

Local development:

```text
http://localhost:8080
```

## Endpoints

### `GET /`

Purpose:
- simple root endpoint

Success response:

- status: `200 OK`
- content type: `text/plain`

Example response body:

```text
Shortens your URL!
```

### `GET /health`

Purpose:
- liveness check

Success response:

- status: `200 OK`
- content type: `text/plain`

Example response body:

```text
Ok!
```

### `GET /ready`

Purpose:
- readiness check

Success response:

- status: `200 OK`
- content type: `text/plain`

Example response body:

```text
Ready!
```

### `POST /shorten`

Purpose:
- create a short code for a URL

Request headers:

- `Content-Type: application/json`

Request body:

```json
{
  "url": "https://example.com"
}
```

Success response:

- status: `200 OK`
- content type: `application/json`

Example response:

```json
{
  "message": "short URL created successfully",
  "url_model": {
    "original_url": "https://example.com",
    "short_code": "100680ad",
    "created_at": "2026-07-28T12:00:00Z",
    "updated_at": "2026-07-28T12:00:00Z",
    "expires_at": null
  }
}
```

Error responses:

- `400 Bad Request`
- `405 Method Not Allowed`
- `500 Internal Server Error`

Example error response:

```json
{
  "error": "bad request"
}
```

Validation rules:

- `url` must be present
- `url` must not be empty after trimming spaces
- `url` must parse as a URL with both scheme and host

### `GET /{code}`

Purpose:
- look up a short code and redirect to the original URL

Path parameter:

- `code`

Success response:

- status: `302 Found`
- header: `Location: <original-url>`

Example:

```text
GET /100680ad
```

Response:

```text
HTTP/1.1 302 Found
Location: https://example.com
```

Error responses:

- `404 Not Found`
- `405 Method Not Allowed`

Example not-found response:

```json
{
  "error": "not found"
}
```

## Response Conventions

JSON endpoints currently return:

- `application/json` for structured success and error responses
- snake_case field names in JSON

Plain-text endpoints currently return:

- `text/plain`

## Notes

- short codes are currently generated deterministically from the original URL
- expired links are treated as not found during resolution
- the API does not yet allow clients to set expiration during creation
