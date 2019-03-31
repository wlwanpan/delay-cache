package dcache

import (
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	dcache := New(time.Millisecond)
	dcache.Set("key1", "val1")
	dcache.Set("key2", "val2")

	var key string
	key = dcache.pop()
	if key != "key1" {
		t.Errorf("Expected 'key1' got: %s", key)
	}
	key = dcache.pop()
	if key != "key2" {
		t.Errorf("Expected 'key2' got: %s", key)
	}

	if len(dcache.queue) != 0 {
		t.Errorf("Expected len of queue to be 0 got: %d", len(dcache.queue))
	}
}

// TestSimpleCRUD simple operations with no worker on.
func TestSimpleCRUD(t *testing.T) {
	dcache := New(time.Second)

	testKey := "testing-key"
	testVal := "testing-val"
	dcache.Set(testKey, testVal)

	if dcache.Size() != 1 {
		t.Errorf("Expected %d got: %d", 1, dcache.Size())
	}

	val, err := dcache.Get(testKey)
	if err != nil {
		t.Error(err)
	}
	if val != testVal {
		t.Errorf("Expected %s got: %s", testVal, val)
	}

	if err := dcache.Remove(testKey, true); err != nil {
		t.Error(err)
	}
	if dcache.Has(testKey) {
		t.Errorf("Expected %s to be deleted", testKey)
	}
}

func TestWorker(t *testing.T) {
	dcache := New(10 * time.Millisecond)
	testVals := []string{"1", "2"}

	// Checking for collected entries
	go func() {
		select {
		case c := <-dcache.Collect():
			// validating value collected
			if !Contain(testVals, c.val.(string)) {
				t.Errorf("Collected invalid value: %s", c.val)
			}
		}
	}()

	dcache.Set("1", "1")
	dcache.Set("2", "2")
	dcache.StartCycle()

	time.Sleep(30 * time.Millisecond)

	if dcache.Size() != 0 {
		t.Errorf("Expected 0 got: %d", dcache.Size())
	}
}

func Contain(arr []string, val string) bool {
	for _, str := range arr {
		if str == val {
			return true
		}
	}
	return false
}
