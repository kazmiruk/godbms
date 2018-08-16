package pool

import (
	"godbms/src/config"
)

type Pool struct {
	queueChan      chan WorkerInterface
	processingChan chan WorkerInterface
}

func NewPool(conf config.Config) *Pool {
	return &Pool{
		queueChan:      make(chan WorkerInterface, conf.MaxWorkers),
		processingChan: make(chan WorkerInterface, conf.MaxWorkers),
	}
}

func (pool *Pool) Add(wrapper WorkerInterface) {
	pool.queueChan <- wrapper
	wrapper.OnAddToPool()
}

func (pool *Pool) ProcessNext() {
	wrapper := <-pool.queueChan
	pool.processingChan <- wrapper
	wrapper.OnGetFromPool()

	go func() {
		wrapper.Run()
		<-pool.processingChan
	}()
}
