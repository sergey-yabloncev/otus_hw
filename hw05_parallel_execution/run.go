package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

type counter struct {
	errCnt int32
}

func (c *counter) add() {
	atomic.AddInt32(&c.errCnt, 1)
}

func (c *counter) load() int32 {
	return atomic.LoadInt32(&c.errCnt)
}

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrNegativeWorkerCount = errors.New("negative worker count")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		m = len(tasks) + 1
	}

	if n <= 0 {
		return ErrNegativeWorkerCount
	}

	var wg sync.WaitGroup
	ctr := counter{}
	tackCh := make(chan Task)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range tackCh {
				if err := task(); err != nil {
					ctr.add()
				}
			}
		}()
	}

	for _, task := range tasks {
		if ctr.load() >= int32(m) {
			break
		}
		tackCh <- task
	}

	close(tackCh)

	wg.Wait()

	if ctr.errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
