# PostgreSQL Local Setup

This document covers the local machine setup needed before running the URL shortener against PostgreSQL.

## Purpose

Use this setup when:

- `psql` is not installed
- PostgreSQL is not running
- the project database does not exist yet
- the migration has not been applied yet

## 1. Check whether Homebrew is installed

Run:

```bash
brew --version
```

What to check:

- if this prints a version, Homebrew is installed
- if `brew` is not found, install Homebrew first from [brew.sh](https://brew.sh)

## 2. Install PostgreSQL

Run:

```bash
brew install postgresql
```

What this installs:

- PostgreSQL server
- `psql`
- `createdb`
- other PostgreSQL command-line tools

What to check:

- command completes successfully

## 3. Start PostgreSQL

Run:

```bash
brew services start postgresql
```

What this does:

- starts PostgreSQL in the background
- makes it easier to restart the DB later

What to check:

- command completes without error

## 4. Verify `psql` is available

Run:

```bash
psql --version
```

What to check:

- you see a PostgreSQL version

If `psql` is still not found:

- close and reopen the terminal
- then run `psql --version` again

## 5. Check that the PostgreSQL server accepts connections

Run:

```bash
psql postgres
```

What to check:

- you enter the PostgreSQL prompt
- it looks like:

```text
postgres=#
```

To exit:

```sql
\q
```

## 6. Create the project database

Run:

```bash
createdb url_shortener
```

If that does not work, run:

```bash
psql postgres -c "CREATE DATABASE url_shortener;"
```

What to check:

- the database is created successfully

## 7. Verify the database exists

Run:

```bash
psql postgres -c "\l"
```

What to check:

- `url_shortener` appears in the database list

## 8. Apply the initial migration

Run this from the project root:

```bash
psql url_shortener -f migrations/000001_create_urls_table.up.sql
```

What this does:

- connects to the `url_shortener` database
- reads SQL from the migration file
- creates the initial schema

What to check:

- command completes without error

## 9. Verify the table exists

Run:

```bash
psql url_shortener -c "\dt"
```

What to check:

- the `urls` table appears

## 10. Inspect the table structure

Run:

```bash
psql url_shortener -c "\d urls"
```

What to check:

- the table contains the expected columns
- you can see constraints such as the unique constraint on `short_code`

## 11. Set `DATABASE_URL` for the app

Run:

```bash
export DATABASE_URL="postgres://localhost:5432/url_shortener?sslmode=disable"
```

Verify:

```bash
echo $DATABASE_URL
```

What to check:

- the value is printed correctly

## 12. Optional: set `PORT` explicitly

Run:

```bash
export PORT=8080
```

This is optional because your app already defaults to `8080`.

## 13. Quick startup check

Run:

```bash
go run ./cmd/server
```

What to check:

- the app starts
- no database connection error appears

Stop the server with:

```text
Ctrl+C
```

## Useful `psql` commands

These are `psql` meta commands, not SQL:

```sql
\l
\dt
\d urls
\q
```

What they do:

- `\l` lists databases
- `\dt` lists tables
- `\d urls` describes the `urls` table
- `\q` exits `psql`

## Common issues

### `brew: command not found`

Meaning:

- Homebrew is not installed

Fix:

- install Homebrew from [brew.sh](https://brew.sh)

### `psql: command not found`

Meaning:

- PostgreSQL is not installed, or shell path is not refreshed yet

Fix:

- install PostgreSQL
- reopen terminal

### Connection refused

Meaning:

- PostgreSQL server is not running

Fix:

```bash
brew services start postgresql
```

### `database "url_shortener" does not exist`

Meaning:

- database creation step was not completed

Fix:

```bash
createdb url_shortener
```

## Completion checklist

Local PostgreSQL setup is complete when all of these are true:

- `psql --version` works
- PostgreSQL is running
- `url_shortener` database exists
- migration file has been applied
- `urls` table exists
- `DATABASE_URL` is set
