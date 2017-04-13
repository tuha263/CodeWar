package main

import "fmt"

const (
	maxnm = 5000
	maxx  = 1000000000
)

var (
	n, m int
	s, f [maxnm + 2]int
)

func main() {

	//read data
	fmt.Scanln(&n, &m)

	for i := 1; i <= n; i++ {
		var (
			ele     int
			useLess float64
		)
		fmt.Scanln(&ele, &useLess)
		s[i] = ele
	}

	//quy hoach dong
	s[n+1] = maxx + 1
	s[0] = 0
	f[n+1] = 1
	for i := n; i >= 0; i-- {
		var max = 0
		for j := i + 1; j <= n+1; j++ {
			if s[j] >= s[i] && f[j] > max {
				max = f[j]
			}
		}
		f[i] = max + 1
	}

	fmt.Print(n - (f[0] - 2))
}
