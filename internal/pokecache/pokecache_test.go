package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Millisecond)
	if c.cache == nil {
		t.Error("The cache map should not be nil")
	}
}

func TestCache_Add(t *testing.T) {
	c := NewCache(time.Millisecond)

	cases := []struct {
		inputKey   string
		inputVal   []byte
		expectedOk bool
	}{
		{
			inputKey: "test",
			inputVal: []byte("test"),
		},
		{
			inputKey: "",
			inputVal: []byte("test"),
		},
	}

	for _, cs := range cases {
		c.Add(cs.inputKey, cs.inputVal)
		if c.cache[cs.inputKey].val == nil {
			t.Error("The value should not be nil")
		}
	}
}

func TestReap(t *testing.T) {
	interval := 10 * time.Millisecond
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("test"))

	time.Sleep(interval + 1*time.Millisecond)

	if _, ok := c.Get(keyOne); ok {
		t.Error("The key should have been reaped")
	}
}

func TestReapFail(t *testing.T) {
	interval := 10 * time.Millisecond
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("test"))

	time.Sleep(1 * time.Millisecond)

	if _, ok := c.Get(keyOne); !ok {
		t.Error("The key should not have been reaped")
	}
}
