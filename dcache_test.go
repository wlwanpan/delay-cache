package dcache

import (
	"testing"
	"time"
)

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
