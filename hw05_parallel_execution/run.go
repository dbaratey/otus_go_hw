package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	// Place your code here
	cntrErr := int32(0)

	wg := &sync.WaitGroup{}

	cntr := 0
	for cntr <= len(tasks) {
		for i := 0; i < n; i++ {
			cntr++
			if cntr <= len(tasks) {
				wg.Add(1)
				go func(pCntrErr *int32, f func() error, wg *sync.WaitGroup) {
					err := f()
					if err != nil {
						atomic.AddInt32(pCntrErr, 1)
					}
					wg.Done()
				}(&cntrErr, tasks[cntr-1], wg)
			}
		}
		wg.Wait()
		if cntrErr >= int32(m) {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
