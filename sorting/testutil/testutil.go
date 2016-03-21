package testutil

import (
	"testing"

	"github.com/seregayoga/shuffle"
)

// AssertSort tests sort function
func AssertSort(t *testing.T, sortFunc func([]int)) {
	a := make([]int, 1000000)
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
