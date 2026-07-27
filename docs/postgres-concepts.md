# PostgreSQL Concepts

## Why We Need PostgreSQL

The current in-memory repository loses all data when the server restarts.
PostgreSQL gives the project durable storage so short links continue working after process restarts.
It also allows multiple app instances to share the same stored data.

## Table Row And Column

A table stores one kind of entity.
For this project, the first table will store shortened URLs.
Each row will represent one shortened URL, and each column will store one field such as short code, original URL, or created time.

## Primary Key

A primary key uniquely identifies each row in the table.
It is common to use a dedicated internal `id` as the primary key.
That keeps row identity stable even if other business fields change later.

## Unique Constraint

`short_code` must be unique.
If two rows had the same short code, redirect behavior would be ambiguous.
The database should enforce this rule, not only the application code.

## Index

The main read query is lookup by short code.
That means the database should support fast search on `short_code`.
A unique constraint on `short_code` also gives us an index for efficient lookup.

## What Queries Matter First

The first schema only needs to support two important queries:
- insert a newly shortened URL
- find the original URL by short code

That means the first schema should optimize for safe writes and fast short-code lookups.
