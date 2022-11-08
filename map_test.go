package iter

import (
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	m := IterMap[string, int](map[string]int{
		"HELLO": 1, "world": 2, "MY": 3, "name": 1, "is": 6, "jordan": 7,
	})

	m2 := MapCollect(Take(Filter[KeyValue[string, int]](m, func(kv KeyValue[string, int]) (include bool) {
		return strings.ToUpper(kv.K) != kv.K
	}), 10), func(el KeyValue[string, int]) (k string, v int) {
		return el.K, el.V
	})

	if len(m2) != 4 {
		t.Errorf("Expected final collected map to have length 4")
	}

	expected := map[string]int{
		"world": 2, "name": 1, "is": 6, "jordan": 7,
	}
	for k2, v2 := range m2 {
		if v, ok := expected[k2]; !ok {
			t.Errorf("Unexpected key %v found in map", k2)
		} else if v != v2 {
			t.Errorf("Unexpected map value at key %v, expected %v but got %v", k2, v, v2)
		} else {
			delete(expected, k2)
		}
	}

	for k, v := range expected {
		t.Errorf("Missing key value pair: (%v, %v)", k, v)
	}
}
