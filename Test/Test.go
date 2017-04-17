package main

import "fmt"

var (
	n   int
	arr [100]int
)

func main() {
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	QuickSort(&arr, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}

//QuickSort sort arr
func QuickSort(arr *[100]int, left, right int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	for l <= r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}

		if l <= r {
			arr[l], arr[r] = Swap(arr[l], arr[r])
			l++
			r--
		}
	}
	if r > left {
		QuickSort(arr, left, r)
	}

	if l < right {
		QuickSort(arr, l, right)
	}
}

//Swap swap value of a and b
func Swap(a, b int) (int, int) {
	return b, a
}
