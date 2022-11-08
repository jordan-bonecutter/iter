package iter

import (
	"testing"
)

func TestSlice(t *testing.T) {
	s := IterSlice[int]([]int{1, 2, 3, 4, 5})
	squared := Collect(Map[int, int](s, func(i int) (o int, stop bool) {
		return i * i, false
	}))

	if len(squared) != 5 {
		t.Errorf("Expected the final collected array to have length == 5, got %v", len(squared))
	}

	if cap(squared) != 5 {
		t.Errorf("Expected the final collected array to have capacity == 5, got %v", cap(squared))
	}

	for idx, el := range squared {
		if el != (idx+1)*(idx+1) {
			t.Errorf("Expected i'th element to be %v, got %v", el, (idx+1)*(idx+1))
		}
	}
}
