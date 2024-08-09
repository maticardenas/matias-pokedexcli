package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	c := NewCache()
	if c.cache == nil {
		t.Error("The cache map should not be nil")
	}
}

func TestCache_Add(t *testing.T) {
	c := NewCache()

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
