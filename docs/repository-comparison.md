# Repository Comparison

## What Stays The Same

Both repositories implement the same repository abstraction.
Both support the two current operations:
- save a shortened URL
- find a URL by short code

This keeps the service layer mostly unchanged.

## What Changes With PostgreSQL

PostgreSQL changes storage from process-local memory to durable shared storage.
Data survives process restarts and can be shared across multiple app instances.
Lookups can use database indexes instead of scanning an in-memory slice.

## Durability

The in-memory repository loses all data when the server restarts.
The PostgreSQL repository persists data beyond process lifetime.
That is one of the biggest differences between the two implementations.

## Uniqueness

The in-memory repository does not currently enforce uniqueness at the storage layer.
The PostgreSQL repository can enforce unique `short_code` values through a database constraint.
That means PostgreSQL can prevent invalid duplicate states more reliably.

## Failure Modes

The in-memory repository has almost no external operational failure modes.
The PostgreSQL repository can fail because of connection issues, database outages, query failures, or scan errors.
This makes the DB-backed version more realistic but also more operationally complex.

## Why The Current Interface Is Starting To Show Limits

The current repository interface does not return `error`.
That was acceptable for the in-memory repository, but it is weak for database-backed storage.
Real database operations can fail for reasons other than not-found, and the current interface cannot express that cleanly.
