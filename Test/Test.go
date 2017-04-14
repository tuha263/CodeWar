package main

import "fmt"

//Ring is info of ring
type Ring struct {
	a, b, h int
}

//RingPoiter point to ring
type RingPoiter struct {
	value   *int
	pointer *Ring
}

var (
	n       int
	ring    [4]Ring
	pointer [8]RingPoiter
)

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		pointer[i*2].pointer = &ring[i]
		pointer[i*2].value = &ring[i].a

		pointer[i*2+1].pointer = &ring[i]
		pointer[i*2+1].value = &ring[i].b
	}

	*pointer[0].value = 1
	pointer[0], pointer[7] = Swap(pointer[0], pointer[7])

	for i := 0; i < n*2; i++ {
		fmt.Println(pointer[i].pointer)
	}

}

//Swap ...
func Swap(a, b RingPoiter) (RingPoiter, RingPoiter) {
	return b, a
}
