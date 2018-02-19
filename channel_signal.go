package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(cat int, wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		if cancelled() {
			fmt.Println("recv signal", cat)
			return
		}

		fmt.Println("job ", cat, "start")
		time.Sleep(5 * time.Millisecond)
	}
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
		
	// select 를 탈출하려면 done이 close 되어야 함
	}
}

var done = make(chan struct{})

func main() {
	//categorySignalMap := make(map[int]chan bool)

	wg := new(sync.WaitGroup) // 대기 그룹 생성


	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	time.Sleep(1 * time.Second)
	done <- struct{}{}
	close(done)

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다림
}
