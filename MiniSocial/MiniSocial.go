package main

import (
	"fmt"
	"math"
)

const (
	//MAXN is max of n
	MAXN int = 1000000
	//MAXV is max of v
	MAXV int = 1000000000
)

var (
	a         [MAXN]int
	f         [MAXN]int64
	n, status int
	result    int64
)

func main() {
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	f[0] = 0
	f[1] = GetDifferent(a[0], a[1])

	for i := 2; i < n; i++ {
		dif := GetDifferent(a[i], a[i-1])
		if (a[i]-a[i-1])*(a[i-1]-a[i-2]) < 0 {
			f[i] = int64(math.Max(float64(f[i-2]+dif), float64(f[i-1])))
		} else {
			f[i] = f[i-1] + dif
		}
	}

	fmt.Println(f[n-1])
}

//GetDifferent of two int
func GetDifferent(a, b int) int64 {
	return int64(math.Abs(float64(a) - float64(b)))
}
