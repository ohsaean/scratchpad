package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sharedValue Processor = Processor{}

type Processor struct {
	// Related
	mu            sync.Mutex
	sum           int
	anotherVar    int
	yetAnotherVar int

	// Not related
	somethingElse int
}

func process(n string) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < 10000; i++ {
			sharedValue.mu.Lock()
			sharedValue.sum = sharedValue.sum + 1
			sharedValue.mu.Unlock()
		}

		fmt.Println("From "+n+":", sharedValue.sum)
	}()
}

func main() {
	processes := []string{"A", "B", "C", "D", "E"}
	for _, p := range processes {
		process(p)
	}

	wg.Wait()
	fmt.Println("Final Sum:", sharedValue.sum)
}
