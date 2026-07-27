# Persistence Notes

## Why In-Memory Storage Is Not Enough

The current repository stores data only inside the running Go process.
If the server restarts, all shortened URLs are lost.
That is fine for early learning, but a real URL shortener must keep links working after restart.

## What Data Must Be Persisted

At minimum, the service must persist:
- short code
- original URL
- created at timestamp

These are the core fields required to shorten a URL once and resolve it later.

## What Storage Guarantees We Need

The storage layer must provide durability so data survives restarts.
It must support fast lookup by short code.
It should enforce uniqueness so two rows do not represent the same short code incorrectly.

## Why PostgreSQL Is A Good Fit

PostgreSQL gives us durable storage, constraints, indexes, and strong relational modeling.
It is common in production systems and teaches concepts that transfer well to other backend services.
It is a good first database for learning correctness, schema design, and migrations.
