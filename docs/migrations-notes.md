# Migration Notes

## What A Migration Is

A migration is a versioned database schema change stored as a file in the repository.
Instead of manually editing database structure, we apply tracked schema changes in order.

## Why We Use `up` And `down`

`up` migrations move the schema forward.
`down` migrations reverse a change so local development and rollback are easier to manage.

## Why Schema Changes Should Be Versioned

Versioned migrations make schema setup repeatable across environments.
They also make database changes reviewable and keep schema evolution history inside the codebase.
