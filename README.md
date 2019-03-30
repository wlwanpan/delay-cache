# Delay cache

[![codecov](https://codecov.io/gh/wlwanpan/delay-cache/branch/master/graph/badge.svg)](https://codecov.io/gh/wlwanpan/delay-cache)
[![GoDoc](https://godoc.org/github.com/wlwanpan/delay-cache?status.svg)](https://godoc.org/github.com/wlwanpan/delay-cache)

A lightweight go package in-memory caching with a queued timeout.
Every entry added to the cache is also "queued" and at a set interval
an entry is popped from the "queue" (removed from the cache).