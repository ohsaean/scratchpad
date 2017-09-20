package main

import (
	"fmt"
	"math"
	"time"
)

func GetPrimes7(n int) int {

	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}

	// optimized
	//res := make([]int, 0, n)
	//s := make([]int, 0, int((n/2)+1))

	var res []int
	var s []int
	for i := 3; i < n+1; i += 2 {
		s = append(s, i)
	}

	mroot := int(math.Sqrt(float64(n)))
	half := len(s)
	i := 0
	m := 3
	for m <= mroot {
		if s[i] != 0 {
			j := int((float64(m*m) - 3.0) * 0.5)
			s[j] = 0
			for j < half {
				s[j] = 0
				j += m
			}
		}
		i++
		m = (2 * i) + 3
	}

	res = append(res, 2)

	//for _, v := range s {
	//	if v != 0 {
	//		res = append(res, v)
	//	}
	//}

	for k := 0; k < len(s); k++ {
		if s[k] != 0 {
			res = append(res, s[k])
		}
	}

	return len(res)
}

func main() {
	max := 10
	var total time.Duration
	for i := 0; i < max; i++ {
		start := time.Now()
		size := GetPrimes7(10000000)
		fmt.Printf("Found %d prime numbers.\n", size)
		elapsed := time.Since(start)
		fmt.Printf("Elapsed Time : %s\n", elapsed)
		total += elapsed
	}
	fmt.Printf("AVG Elapsed Time : %s\n", total/time.Duration(max))
}
