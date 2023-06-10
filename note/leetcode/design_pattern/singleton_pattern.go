package design_pattern

import "sync"

type pool struct {
	Count int
}

var (
	poolObj *pool
	once    sync.Once
	mu      sync.Mutex
)

func SinglePool() *pool {
	once.Do(func() {
		poolObj = &pool{Count: 10}
	})
	return poolObj
}

func SinglePool1() *pool {
	mu.Lock()
	if poolObj == nil {
		poolObj = &pool{Count: 10}
	}
	mu.Unlock()
	return poolObj
}
