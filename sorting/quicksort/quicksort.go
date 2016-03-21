package quicksort

// Sort orders the input slice in ascending order
func Sort(a []int) {
	if len(a) <= 1 {
		return
	}

	i := partition(a)
	Sort(a[:i])
	Sort(a[i+1:])
}

func partition(a []int) int {
	lastIndex := len(a) - 1
	mid := a[lastIndex]

	i := 0
	j := lastIndex - 1
	for {
		for a[i] < mid {
			i++
		}

		for a[j] > mid {
			if j == 0 {
				break
			}
			j--
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}

	a[i], a[lastIndex] = a[lastIndex], a[i]

	return i
}
