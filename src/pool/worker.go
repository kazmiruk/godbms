package pool

type WorkerInterface interface {
	Run()
	OnAddToPool()
	OnGetFromPool()
}
