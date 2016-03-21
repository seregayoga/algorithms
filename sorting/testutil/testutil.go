package testutil

import (
	"testing"

	"github.com/seregayoga/shuffle"
)

// AssertSort tests sort function
func AssertSort(t *testing.T, sortFunc func([]int), size int) {
	a := make([]int, size)
	for i := range a {
		a[i] = i
	}
	shuffle.Shuffle(a)

	sortFunc(a)

	prev := a[0]
	for _, item := range a {
		if item < prev {
			t.FailNow()
		}

		prev = item
	}
}
