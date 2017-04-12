package main

import "fmt"

func main() {
	var n, a, b, c, r1, r2 int64
	fmt.Scanln(&n)
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	r1 = n / a
	r2 = 0
	if n >= b {
		r2++
	}
	r2 += (n - b) / (b - c)
	r2 += (n - r2*(b-c)) / a
	result := r2
	if r1 > r2 {
		result = r1
	}
	fmt.Println(result)
}
