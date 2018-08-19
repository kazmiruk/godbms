package storages

import (
	"godbms/src/config"
	"godbms/src/structures"
	"godbms/src/structures/kvstorage"
)

type Storages struct {
	conf     config.Config
	storages map[string]structures.Storage
}

func NewStorages(conf config.Config) (storages *Storages) {
	return &Storages{
		conf,
		make(map[string]structures.Storage),
	}
}

func (storages *Storages) GetStorage(name string) (storage structures.Storage) {
	return storages.storages[name]
}

func (storages *Storages) CreateStorage(t string, name string) {
	var storage structures.Storage

	switch t {
	case "kv":
		storage = kvstorage.NewKeyValueStorage(name, storages.conf.GarbageCollectorTick)
	}

	storages.storages[name] = storage
}
