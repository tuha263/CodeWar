package main

import (
	"fmt"
	"math"
)

const (
	maxnm = 5000
	maxx  = 1000000000
)

var n, m int
var s, count [maxnm]int
var f [maxnm][maxnm]int

func main() {

	//read data
	fmt.Scanln(&n, &m)

	for i := 0; i < n; i++ {
		var ele int
		fmt.Scanln(&ele)
		s[i] = ele
		count[ele]++
	}

	//int
	for i := 0; i < m; i++ {
		f[0][i] = 1
	}

	for i := 0; i < n; i++ {
		var k int
		if s[0] == 1 {
			k = 1
		} else {
			k = 0
		}
		f[i][0] = k
	}

	//quy hoach dong
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			f[i][j] = int(math.Max(float64(f[i-1][j]+1), float64(f[i][j-1])))
		}
	}
	if s[0] > 1 {
		fmt.Print(1)
	}
	//	fmt.Print(n - f[n][m])
}
