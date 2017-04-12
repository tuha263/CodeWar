package main

import "fmt"

const (
	maxnm = 5000
	maxx  = 1000000000
)

var n, m int
var s [maxnm + 1]int
var f [maxnm + 1]int

func main() {

	//read data
	fmt.Scanln(&n, &m)

	for i := 0; i < n; i++ {
		var ele int
		fmt.Scanln(&ele)
		s[i] = ele
	}

	//quy hoach dong
	s[n] = maxx + 1
	f[n] = 1
	for i := n - 1; i >= 0; i-- {
		var max = 0
		for j := i + 1; j <= n; j++ {
			if s[j] >= s[i] && f[j] > max {
				max = f[j]
			}
		}
		f[i] = max + 1
	}

	var result int
	for i := 0; i < n; i++ {
		if result < f[i] {
			result = f[i]
		}
	}

	fmt.Print(n - result + 1)
}
