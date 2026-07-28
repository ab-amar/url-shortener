# Expiration Notes

## What Expiration Means

Expiration means a short link should stop being considered valid after a certain time.
The row may still exist in storage, but the application should no longer serve redirects for it once it has expired.

## Why We Use `expires_at`

`expires_at` stores the exact timestamp after which the link is no longer valid.
This makes expiry checks simple because the application can compare the current time with a single stored timestamp.

## Why `expires_at` Should Be Optional

Not every short link needs an expiration time.
Some links may live forever, so the field should be optional.
In Go, `*time.Time` is a good fit because `nil` can mean “no expiration”.

## Why A Timestamp Is Better Than A Boolean

A boolean only tells whether the link is expired right now.
A timestamp tells both whether it is expired and when it will expire.
That makes it much more useful for debugging, operations, and future product behavior.
