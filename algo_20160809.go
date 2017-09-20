package main

import (
	. "github.com/y0ssar1an/q"
	"strconv"
)

func main() {

	m := make(map[string]int)

	for i := 1; i <= 1000; i++ {
		str := strconv.Itoa(i)
		for _, v := range str {
			s := string(v)
			m[s]++
		}
	}
	//fmt.Println(m)
	Q("dss")
	// map[0:192 2:300 3:300 5:300 7:300 9:300 1:301 4:300 6:300 8:300]
}
