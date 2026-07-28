# Constraints And Indexes

## Primary Key

The primary key is the main internal identity for each row in the table.
It must be unique and cannot be null.
For this project, the `id` column should be the primary key.

## Why `id` Is The Primary Key

`id` is a stable internal identifier for database rows.
`short_code` is a business-facing value used in redirects, but `id` is better for internal row identity.
Keeping internal identity separate from public identity makes the schema easier to evolve later.

## Unique Constraint On `short_code`

`short_code` must be unique because one code must map to exactly one destination URL.
If duplicates were allowed, redirect behavior would become ambiguous and incorrect.
The database should enforce this rule instead of relying only on application logic.

## Why `short_code` Needs Fast Lookup

The main read query in this project is: find a row by short code.
That means lookups on `short_code` should be efficient.
In PostgreSQL, a unique constraint on `short_code` also gives us an index for fast lookups.

## What We Are Not Indexing Yet

We are not indexing `original_url` or `created_at` yet.
The current application behavior does not query by those columns.
Indexes should be added based on real query patterns, not just because a column exists.
