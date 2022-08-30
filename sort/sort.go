package sort

func bubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
}
func selectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[i] {
				ar[j], ar[i] = ar[i], ar[j]
			}
		}
	}
}

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	a := mergeSort(ar[:len(ar)/2])
	b := mergeSort(ar[len(ar)/2:])
	var res []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return res
}

func quickSort(ar []int, s, f int) {
	if s < f {
		p := ar[f]
		i := s
		for j := s; j < f; j++ {
			if ar[j] < p {
				ar[i], ar[j] = ar[j], ar[i]
				i++
			}
		}
		ar[i], ar[f] = ar[f], ar[i]

		quickSort(ar, s, i-1)
		quickSort(ar, i+1, f)
	}
}

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		for j := 0; j < i; j++ {
			if ar[j] > ar[i] {
				ar[j], ar[i] = ar[i], ar[j]
			}
		}
	}
}
