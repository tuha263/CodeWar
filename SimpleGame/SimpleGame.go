package main

import (
	"fmt"
	"math"
)

const (
	maxLen  = 64
	maxAlen = 200000
)

var countArr [maxLen]int

func main() {
	var n, m, k int32
	var a [maxAlen]int32
	var max int64
	n = 0
	fmt.Scan(&n, &m, &k)
	powValue := int64(math.Pow(float64(k), float64(m)))

	for i := int32(0); i < n; i++ {
		var ele int32
		fmt.Scan(&ele)
		a[i] = ele
		countArr = CountBit(int64(ele), true, countArr)
	}

	for i := int32(0); i < n; i++ {
		arr1 := CountBit(int64(a[i]), false, countArr)
		arr1 = CountBit(int64(a[i])*powValue, true, arr1)
		var value = GetValue(arr1)
		if value > max {
			max = value
		}
	}

	fmt.Print(max)
}

//CountBit add cout of bit
func CountBit(a int64, isAdd bool, countArr [maxLen]int) [maxLen]int {
	var index int
	for a > 0 {
		if a&1 == 1 {
			if isAdd {
				countArr[index]++
			} else {
				countArr[index]--
			}
		}
		a >>= 1
		index++
	}

	return countArr
}

//GetValue return value of list bit
func GetValue(countArr [maxLen]int) int64 {
	var result int64
	for i := 0; i < maxLen; i++ {
		if countArr[i] > 0 {
			result += int64(math.Pow(2, float64(i)))
		}
	}
	return result
}
