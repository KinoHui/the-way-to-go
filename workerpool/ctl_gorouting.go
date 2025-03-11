package workerpool

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func biz(ch chan int) {
	for i := range ch {
		fmt.Println("go func ", i, " goroutine count: ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(ch chan int, i int) {
	wg.Add(1)
	ch <- i
}

func Start() {
	ch := make(chan int)

	goCnt := 10

	for i := 0; i < goCnt; i++ {
		go biz(ch)
	}

	task_num := 100000

	for i := 0; i < task_num; i++ {
		sendTask(ch, i)
	}

	wg.Wait()
}
