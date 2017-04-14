package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Ring is info of ring
type Ring struct {
	a, b, h int
}

const (
	//MAXN is max of n
	MAXN = 100000
	//MAXX is max of a,b and h
	MAXX = 1000000000
)

var (
	n    int
	ring [MAXN + 1]Ring
	f    [MAXN + 1]int64
)

func main() {
	fmt.Scan(&n)

	in := bufio.NewReader(os.Stdin)

	ss, _ := in.ReadString('.')

	ls := strings.Split(ss, string('\n'))
	for i := 0; i < n; i++ {
		abh := strings.Split(ls[i], " ")
		a, _ := strconv.Atoi(abh[0])
		b, _ := strconv.Atoi(abh[1])
		h, _ := strconv.Atoi(abh[2])
		ring[i] = Ring{a, b, h}
	}

	QuickSort(0, n-1)
	ring[n] = Ring{0, MAXX + 1, 0}

	f[0] = int64(ring[0].h)
	for i := 0; i <= n; i++ {
		var max int64
		for j := 0; j < i; j++ {
			if ring[i].b > ring[j].a && max < f[j] {
				max = f[j]
			}
			f[i] = max + int64(ring[i].h)
		}
	}

	fmt.Println(f[n])
}

//QuickSort sort ring list
func QuickSort(left, right int) {
	l := left
	r := right
	pivot := ring[(left+right)/2].b
	for l < r {
		for ring[l].b > pivot {
			l++
		}

		for ring[r].b < pivot {
			r--
		}

		if l <= r {
			temp := ring[l]
			ring[l] = ring[r]
			ring[r] = temp
			l++
			r--
		}

		if left < r {
			QuickSort(left, r)
		}

		if right > l {
			QuickSort(l, right)
		}
	}

}
