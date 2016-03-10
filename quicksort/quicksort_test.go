package quicksort

import (
	"testing"

	"github.com/seregayoga/shuffle"
)

func TestSort(t *testing.T) {
	a := make([]int, 1000000)
	for i := range a {
		a[i] = i
	}
	shuffle.Shuffle(a)

	Sort(a)

	prev := a[0]
	for _, item := range a {
		if item < prev {
			t.FailNow()
		}

		prev = item
	}
}
