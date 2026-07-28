# Timestamps

## `created_at`

`created_at` stores when a row was first created.
It is useful for debugging, audits, analytics, and later expiry logic.
This value should not change after the row is inserted.

## `updated_at`

`updated_at` stores when a row was last modified.
It becomes useful once records can be edited, expired, soft-deleted, or otherwise changed later.
Even if updates are minimal today, it is a realistic production field to keep.

## Why We Keep Both

`created_at` and `updated_at` answer different questions.
One tells us when the record first came into existence, and the other tells us when it was last changed.
Keeping both makes future data operations and debugging easier.

## How They Start On Insert

When a row is first inserted, `created_at` and `updated_at` should start with the same value.
That is because the creation moment is also the last update moment at insert time.
In Go code, it is better to call `time.Now()` once and reuse the same value for both fields.
