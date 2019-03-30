# Delay Cache

A lightweight go package in-memory caching with a queued timeout.
Every entry added to the cache is also "queued" and at a set interval
an entry is popped from the "queue" (removed from the cache).