package testutil

import (
	"sort"
	"testing"
)

func TestAssertSort(t *testing.T) {
	AssertSort(t, sort.Ints, 10000)
}
