package chunky

import "testing"

func TestIterator(t *testing.T) {

	a, err := NewIterator([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	if err != nil {
		t.Error(err)
	}

	t.Log(a)
}
