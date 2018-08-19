package pool

import (
	"godbms/src/config"
)

type Pool struct {
	queueChan chan WorkerInterface
}

func NewPool(conf config.Config) *Pool {
	return &Pool{
		queueChan: make(chan WorkerInterface, conf.MaxWorkers),
	}
}

func (pool *Pool) AddToProcess(wrapper WorkerInterface) {
	wrapper.OnAddToPool()
	pool.queueChan <- wrapper

	go func() {
		wrapper.OnGetFromPool()
		wrapper.Run()
		<-pool.queueChan
	}()
}
