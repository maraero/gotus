package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ch := make(chan Task)
	wg := sync.WaitGroup{}
	errCnt := newErrCnt(m)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					errCnt.inc()
				}
			}
		}()
	}

	for _, task := range tasks {
		if errCnt.exceedsLimit() {
			break
		}
		ch <- task
	}

	close(ch)
	wg.Wait()

	if errCnt.exceedsLimit() {
		return ErrErrorsLimitExceeded
	}

	return nil
}
