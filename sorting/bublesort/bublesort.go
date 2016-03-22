package bubblesort

// Sort orders the input slice in ascending order
func Sort(a []int) {
	if len(a) <= 1 {
		return
	}

	for i := 1; i < len(a); i++ {
		for j := len(a) - 1; j >= i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
