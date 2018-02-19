package main

import (
	"sync"
	"fmt"
)

func job (i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("job ", i)
}


func main() {
	//var wg sync.WaitGroup

	wg := new(sync.WaitGroup)

	for i:=1; i<= 10; i++ {
		wg.Add(1)
		go job(i, wg)
	}

	wg.Wait()
}
