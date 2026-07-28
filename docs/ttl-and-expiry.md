# TTL And Expiry

## What TTL Means

TTL means time to live.
It is usually expressed as a duration such as 1 hour, 24 hours, or 7 days.
TTL is a useful way to think about expiry when a link is first created.

## Why We Store `expires_at` Instead Of TTL

TTL is relative, but `expires_at` is an absolute timestamp.
Absolute timestamps are easier to store, compare, and reason about during lookups.
So TTL is useful as an input idea, while `expires_at` is better as the stored value.

## When A Link Is Expired

If `ExpiresAt == nil`, the link has no expiration and stays active.
If `ExpiresAt` has a value and the current time is after or equal to it, the link is expired.
That means the redirect should no longer be served.

## Expired Does Not Mean Deleted

An expired link can still exist in storage.
Expiry affects whether the link is valid for redirect behavior, while deletion affects whether the row exists at all.
Keeping expired rows can still be useful for debugging, analytics, and later cleanup jobs.

## Expected Lookup Behavior

When resolving a short code, the service should distinguish three states:
- not found
- found but expired
- found and active

Only the active case should redirect normally.
