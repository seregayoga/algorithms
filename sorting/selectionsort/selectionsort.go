package selectionsort

// Sort orders the input slice in ascending order
func Sort(a []int) {
	if len(a) <= 1 {
		return
	}

	for i := 0; i < (len(a) - 1); i++ {
		min := i

		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}

		a[i], a[min] = a[min], a[i]
	}
}
