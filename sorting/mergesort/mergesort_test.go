package mergesort

import (
	"testing"

	"github.com/seregayoga/algorithms/sorting/testutil"
)

func TestSort(t *testing.T) {
	testutil.AssertSort(t, Sort)
}
