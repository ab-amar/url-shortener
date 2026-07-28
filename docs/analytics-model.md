# Analytics Model

## What We Want To Measure

The first analytics goal is to measure how often shortened links are actually used.
For this project, the most important event is a redirect being triggered.
That tells us whether a short link is getting traffic.

## Why We Use An Event Model

An event model stores one record per occurrence.
This is better than only storing a single running counter because it preserves history.
Later, counts and trends can be derived from raw events.

## First Redirect Event Shape

The first redirect event can stay very small:
- `id`
- `short_code`
- `created_at`

This is enough to know which short code was used and when the redirect happened.

## Why We Keep The First Version Small

We do not need to add every possible analytics field in version 1.
Starting small keeps the design easier to implement and reason about.
We can still get useful insight from just short code and timestamp.

## How This Can Grow Later

Later, the event model can expand with fields such as:
- user agent
- referrer
- country
- IP hash

We can also derive aggregate counters and reports from the raw redirect events later.
