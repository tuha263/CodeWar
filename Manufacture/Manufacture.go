package main

import (
	"fmt"
	"math"
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
	f        [MAXN + 1]int64
	pointer  [MAXN*2 + 2]RingPoiter
)

func main() {
	fmt.Scan(&n)

	/*in := bufio.NewReader(os.Stdin)

	ss, _ := in.ReadString('.')

	ls := strings.Split(ss, string('\n'))
	for i := 0; i < n; i++ {
		abh := strings.Split(ls[i], " ")
		a, _ := strconv.Atoi(abh[0])
		b, _ := strconv.Atoi(abh[1])
		h, _ := strconv.Atoi(abh[2])
		ring[i] = Ring{a, b, h}
	}*/
	for i := 0; i < n; i++ {
		fmt.Scan(&ring[i].a, &ring[i].b, &ring[i].h)
	}

	QuickSort(0, n-1)
	//Discrete()

	/*	for i := 0; i < n; i++ {
			f[i] = GetMAX(ring[i].b-1) + int64(ring[i].h)
			UpdateBIT(ring[i].a, f[i])
		}

		fmt.Println(GetMAX(n * 2))*/

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

// UpdateBIT update BIT
func UpdateBIT(i int, value int64) {
	if i <= n*2+1 {
		f[i] = int64(math.Max(float64(value), float64(f[i])))
		UpdateBIT(i+i&(-i), value)
	}
}

// GetMAX get max form 0..i
func GetMAX(i int) int64 {
	//fmt.Println(i)
	if i == 0 {
		return 0
	}
	if i-i&(-i) == 0 {
		return f[i]
	}
	return int64(math.Max(float64(f[i]), float64(GetMAX(i-i&(-i)))))
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
	for i := 0; i < n*2; i++ {
		pointer[i*2].pointer = &ring[i]
		pointer[i*2].value = &ring[i].a

		pointer[i*2+1].pointer = &ring[i]
		pointer[i*2+1].value = &ring[i].b
	}

	QuickSortPoint(0, n*2-1)

	/*for i := 0; i < n*2; i++ {
		fmt.Println(*pointer[i].value)
	}*/

	for i := 0; i < n*2; i++ {
		*pointer[i].value = i + 1
	}
}

//QuickSortPoint sort ring point list
func QuickSortPoint(left, right int) {
	l := left
	r := right
	m := (left + right) / 2
	pivot := *pointer[m].value
	for l < r {
		for *pointer[l].value < pivot || (*pointer[l].value == pivot && l%2 == 1 && m%2 == 0) {
			l++
		}

		for *pointer[r].value > pivot || (*pointer[r].value == pivot && r%2 == 0 && m%2 == 1) {
			r--
		}

		if l <= r {
			/*fmt.Print(l)
			fmt.Print("-")
			fmt.Println(r)*/
			temp := pointer[l]
			pointer[l] = pointer[r]
			pointer[r] = temp
			l++
			r--
		}

		if left < r {
			QuickSortPoint(left, r)
		}

		if right > l {
			QuickSortPoint(l, right)
		}
	}

}
