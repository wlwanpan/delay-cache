# Delay cache

[![codecov](https://codecov.io/gh/wlwanpan/delay-cache/branch/master/graph/badge.svg)](https://codecov.io/gh/wlwanpan/delay-cache)
[![GoDoc](https://godoc.org/github.com/wlwanpan/delay-cache?status.svg)](https://godoc.org/github.com/wlwanpan/delay-cache)

A lightweight go package in-memory caching with a queued (FIFO) timeout.
Every entry added to the cache is also "queued" and at a set interval
an entry is popped from the queue.

When an entry is timed out (popped out of the queue), it can be collected
in a read-only channel.

## Installation
```bash
go get github.com/wlwanpan/delay-cache
```

- Basic example
```go
dcache = dcache.New(5 * time.Second)

err := dcache.Set("key", "val")
...

val, err := dcache.Get("key")
...

has := dcache.Has("key")
...
```