package concurrent

import (
	"fmt"
	"sync"
)

var a int = 0
var wait sync.WaitGroup

func concur() {

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	wait.Add(2)
	go func() {
		for a < 99 {
			<-ch1
			a += 1
			fmt.Println("协程1", a)
			ch2 <- struct{}{}
		}
		wait.Done()
	}()
	go func() {
		for a < 99 {
			<-ch2
			a += 1
			fmt.Println("协程2", a)
			ch1 <- struct{}{}
		}
		wait.Done()
	}()
	ch1 <- struct{}{}
	wait.Wait()
	fmt.Println("exit...")

}
