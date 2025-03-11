package workerpool

import (
	"fmt"
	"time"
)

type WorkerManager struct {
	workerChan chan *worker
	nWorkers   int
}

type worker struct {
	id  int
	err error
}

func NewWorkerManager(nWorkers int) *WorkerManager {
	return &WorkerManager{
		workerChan: make(chan *worker, nWorkers),
		nWorkers:   nWorkers,
	}
}

func (wm *WorkerManager) StartWorkerPool() {
	for i := 0; i < wm.nWorkers; i++ {
		wk := &worker{
			id: i,
		}
		go wk.work(wm.workerChan)
	}

	wm.KeepLivaWorkers()
}

func (wm *WorkerManager) KeepLivaWorkers() {
	for wk := range wm.workerChan {
		fmt.Println("worker", wk.id, "is dead", "[", wk.err, "]")

		wk.err = nil

		go wk.work(wm.workerChan)
	}
}

func (wk *worker) work(workerChan chan *worker) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("worker %d panic: %v", wk.id, r)
			}
		} else {
			wk.err = err
		}

		workerChan <- wk
	}()

	fmt.Println("start work...ID:", wk.id)

	time.Sleep(time.Second * 3)

	panic("worker panic")

	return err
}
