package mergesort

// Sort orders the input slice in ascending order
func Sort(a []int) {
	if len(a) <= 1 {
		return
	}

	mid := len(a) / 2

	Sort(a[:mid])
	Sort(a[mid:])

	merge(a)
}

func merge(a []int) {
	mid := len(a) / 2

	firstPart := a[:mid]
	secondPart := a[mid:]

	mergeResult := make([]int, len(a))

	i := 0
	j := 0
	for k, _ := range a {
		if i == len(firstPart) {
			mergeResult[k] = secondPart[j]
			j++
			continue
		}

		if j == len(secondPart) {
			mergeResult[k] = firstPart[i]
			i++
			continue
		}

		if firstPart[i] < secondPart[j] {
			mergeResult[k] = firstPart[i]
			i++
		} else {
			mergeResult[k] = secondPart[j]
			j++
		}
	}

	copy(a, mergeResult)
}
