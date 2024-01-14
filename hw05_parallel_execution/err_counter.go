package hw05parallelexecution

import "sync/atomic"

type errCnt struct {
	cnt   int32
	limit int32
}

func (c *errCnt) inc() {
	atomic.AddInt32(&c.cnt, 1)
}

func (c *errCnt) exceedsLimit() bool {
	if c.limit <= 0 {
		return false
	}
	return atomic.LoadInt32(&c.cnt) >= c.limit
}

func newErrCnt(limit int) *errCnt {
	return &errCnt{limit: int32(limit)}
}
