package shellsort

// Sort orders the input slice in ascending order
func Sort(a []int) {
	if len(a) <= 1 {
		return
	}

	for i := 1; i < len(a); i++ {
		v := a[i]
		j := i

		for j > 0 && v < a[j-1] {
			a[j] = a[j-1]
			j--
		}

		a[j] = v
	}
}

