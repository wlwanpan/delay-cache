# Delay cache

[![Build Status](https://travis-ci.org/wlwanpan/delay-cache.svg?branch=master)](https://travis-ci.org/wlwanpan/delay-cache)
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

## Usage

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

- Collecting delayed entry
```go
dcache = dcache.New(5 * time.Second)

dcache.Set("key1", "val1")
dcache.Set("key2", "val2")
dcache.Set("key3", "val3")

dcache.StartCycle()

for {
    select {
        case entry := <- dcache.Collect():
            // 1st select to be collected: &Collect{key: "key1", val: "val1"}
            // 2nd select after 5 seconds: &Collect{key: "key2", val: "val3"}
            // ...
    }
}

```
