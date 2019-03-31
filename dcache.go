package dcache

import (
	"errors"
	"time"
)

var (
	// ErrKeyAlreadyInUse is returned when setting an entry on an exist key.
	ErrKeyAlreadyInUse = errors.New("key already in use")

	// ErrKeyDoestNotExist is returned when retrieving a non existant key.
	ErrKeyDoestNotExist = errors.New("key does not exist")
)

type DCache struct {
	tick time.Duration

	entries map[string]interface{}

	queue []string

	out chan *Collect

	stopWorker chan bool
}

type Collect struct {
	key string

	val interface{}
}

func New(tick time.Duration) *DCache {
	if tick == 0 {
		tick = time.Second
	}

	return &DCache{
		tick:       tick,
		entries:    make(map[string]interface{}),
		queue:      []string{},
		out:        make(chan *Collect),
		stopWorker: make(chan bool),
	}
}

func (c *DCache) push(key string) {
	c.queue = append(c.queue, key)
}

func (c *DCache) pop() string {
	if len(c.queue) < 1 {
		return ""
	}

	key := c.queue[0]
	c.queue = c.queue[1:]
	return key
}

func (c *DCache) startWorker() {
	tick := time.Tick(c.tick)
	for {
		select {
		case <-tick:
			key := c.pop()
			entry, err := c.Get(key)
			if err != nil {
				continue
			}

			delete(c.entries, key)
			c.out <- &Collect{
				key: key,
				val: entry,
			}
		case <-c.stopWorker:
			return
		}
	}
}

// StartCycle starts the worker in a go routine.
func (c *DCache) StartCycle() {
	go c.startWorker()
}

// StopCycle signals any worker to stop.
func (c *DCache) StopCycle() {
	c.stopWorker <- true
}

// Collect delayed entries removed from cache.
func (c *DCache) Collect() <-chan *Collect {
	return c.out
}

// Get reads an entry from cache.
func (c *DCache) Get(key string) (interface{}, error) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, ErrKeyDoestNotExist
	}
	return entry, nil
}

// Set saves an entry to the cache.
func (c *DCache) Set(key string, val interface{}) error {
	if has := c.Has(key); has {
		return ErrKeyAlreadyInUse
	}
	c.entries[key] = val
	c.push(key)
	return nil
}

// Has is wrapper around get to check if an entry still exist in the cache.
func (c *DCache) Has(key string) bool {
	_, err := c.Get(key)
	if err != ErrKeyDoestNotExist {
		return true
	}
	return false
}

// Remove deletes an entry from the cache and:
//   silent false: send the val to the out channel to collect.
//   silent  true: val cannot be collected.
func (c *DCache) Remove(key string, silent bool) error {
	entry, err := c.Get(key)
	if err != nil {
		return err
	}
	if !silent {
		c.out <- &Collect{
			key: key,
			val: entry,
		}
	}
	delete(c.entries, key)
	return nil
}

// Size returns the number of entries currently stored.
func (c *DCache) Size() int {
	return len(c.entries)
}
