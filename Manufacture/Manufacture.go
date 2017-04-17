package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Ring is info of ring
type Ring struct {
	a, b, h int
}

//RingPoiter point to ring
type RingPoiter struct {
	value   *int
	pointer *Ring
}

const (
	//MAXN is max of n
	MAXN = 100000
	//MAXX is max of a,b and h
	MAXX = 1000000000
)

var (
	n, maxab int
	ring     [MAXN + 1]Ring
	f        [MAXN*2 + 1]int64
	pointer  [MAXN*2 + 2]RingPoiter
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

	Discrete()
	QuickSort(0, n-1)

	for i := 0; i < n; i++ {
		f[ring[i].a] = GetMAX(ring[i].b-1) + int64(ring[i].h)
		UpdateBIT(ring[i].a, f[ring[i].a])
	}

	fmt.Println(GetMAX(maxab))
}

// UpdateBIT update BIT
func UpdateBIT(i int, value int64) {
	i += i & -i
	if i <= maxab {
		f[i] = int64(math.Max(float64(value), float64(f[i])))
		UpdateBIT(i, value)
	}
}

// GetMAX get max form 0..i
func GetMAX(i int) int64 {
	if i == 0 {
		return 0
	}

	return int64(math.Max(float64(f[i]), float64(GetMAX(i-i&(-i)))))
}

//QuickSort sort ring list
func QuickSort(left, right int) {
	l := left
	r := right
	pivot := ring[(left+right)/2]
	for l <= r {
		for ring[l].b > pivot.b || (ring[l].b == pivot.b && ring[l].a > pivot.a) {
			l++
		}

		for ring[r].b < pivot.b || (ring[r].b == pivot.b && ring[r].a < pivot.a) {
			r--
		}

		if l <= r {
			temp := ring[l]
			ring[l] = ring[r]
			ring[r] = temp
			l++
			r--
		}
	}

	if left < r {
		QuickSort(left, r)
	}

	if right > l {
		QuickSort(l, right)
	}

}

//Discrete discrete a and b of ring
func Discrete() {
	for i := 0; i < n; i++ {
		pointer[i*2].pointer = &ring[i]
		pointer[i*2].value = &ring[i].a

		pointer[i*2+1].pointer = &ring[i]
		pointer[i*2+1].value = &ring[i].b
	}

	QuickSortPoint(0, n*2-1)

	curValue := *pointer[0].value
	maxab = 1
	for i := 0; i < n*2; i++ {
		if *pointer[i].value != curValue {
			maxab++
			curValue = *pointer[i].value
		}
		*pointer[i].value = maxab
	}
}

//QuickSortPoint sort ring point list
func QuickSortPoint(left, right int) {
	l := left
	r := right
	m := (left + right) / 2
	pivot := *pointer[m].value
	for l < r {
		for *pointer[l].value < pivot {
			l++
		}

		for *pointer[r].value > pivot {
			r--
		}

		if l <= r {
			temp := pointer[l]
			pointer[l] = pointer[r]
			pointer[r] = temp
			l++
			r--
		}
	}

	if left < r {
		QuickSortPoint(left, r)
	}

	if right > l {
		QuickSortPoint(l, right)
	}

}
