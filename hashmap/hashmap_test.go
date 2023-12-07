package hashmap

import (
	"hash/fnv"
	"testing"
)

func TestHashing(t *testing.T) {
	hasher := NewHasher(fnv.New32a())

	val1, err := hasher.Calculate("hello")
	if err != nil {
		t.Errorf("Error calculating hash: %v", err)
	}

	val2, err := hasher.Calculate("hello")
	if err != nil {
		t.Errorf("Error calculating hash: %v", err)
	}

	if val1 != val2 {
		t.Errorf("Expected hash of same strings to be same")
	}

	val3, err := hasher.Calculate(123.5423)
	if err != nil {
		t.Errorf("Error calculating hash: %v", err)
	}

	if val3 == val1 {
		t.Errorf("Expected hash of different values to be different")
	}
}

func tassert(t *testing.T, exp bool, msg string, args ...any) {
	if !exp {
		t.Errorf(msg, args...)
	}
}

func TestMap(t *testing.T) {
	mapa := New[string]()
	mapa.InsertOrUpdate("key1", "val1")
	mapa.InsertOrUpdate("key2", "val2")

	if !mapa.Contains("key1") {
		t.Errorf("expected map to contain key1")
	}

	tassert(t, mapa.Contains("key2"), "mapa should contain key2")

	if *mapa.Get("key2") != "val2" {
		t.Errorf("expected value at key2 to be val2, got: %s", *mapa.Get("key2"))
	}

	*mapa.Get("key2") = "new val2"
	if *mapa.Get("key2") != "new val2" {
		t.Errorf("expected value at key2 to be new val2, got: %s", *mapa.Get("key2"))
	}

	if !mapa.Contains("key1") {
		t.Errorf("map does have node with key key1")
	}

	mapa.InsertOrUpdate("key1", "new val1")

	if *mapa.Get("key1") != "new val1" {
		t.Errorf("expected value at key1 to be new val1, got: %s", *mapa.Get("key1"))
	}
}
