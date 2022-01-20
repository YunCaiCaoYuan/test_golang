package wg_wf

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	wg    sync.WaitGroup //将循环栅栏替换成WaitGroup
}

func New() *H2O {
	var wg sync.WaitGroup
	wg.Add(3)

	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		wg:    wg,
	}
}

func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)
	releaseHydrogen()

	// 标记自己已达到，等待其它goroutine到达
	h2o.wg.Done()
	h2o.wg.Wait()

	h2o.semaH.Release(1)
}

func (h2o *H2O) Oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)
	releaseOxygen()

	// 标记自己已达到，等待其它goroutine到达
	h2o.wg.Done()
	h2o.wg.Wait()
	//都到达后重置wg
	h2o.wg.Add(3)

	h2o.semaO.Release(1)
}
