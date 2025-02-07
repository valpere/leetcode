package sort

import "golang.org/x/exp/constraints"

type sortable interface {
	constraints.Ordered | byte
}

func QuickSortInt[T sortable](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

// Sorts (a portion of) an array, divides it into partitions, then sorts those
func quickSort[T sortable](arr []T, lo, hi int) {
	if (lo >= hi) || (lo < 0) {
		return
	}

	lt, gt := partition(arr, lo, hi)

	quickSort(arr, lo, lt-1)
	quickSort(arr, gt+1, hi)
}

// Divides array into three partitions
func partition[T sortable](arr []T, lo, hi int) (lt, gt int) {
	pivot := arr[lo+(hi-lo)/2] // Choose the middle element as the pivot (integer division)

	lt = lo
	gt = hi

	for eq := lo; eq <= gt; {
		if arr[eq] < pivot {
			arr[eq], arr[lt] = arr[lt], arr[eq]
			lt += 1
			eq += 1
		} else if arr[eq] > pivot {
			arr[eq], arr[gt] = arr[gt], arr[eq]
			gt -= 1
		} else {
			eq += 1
		}
	}

	return lt, gt
}
