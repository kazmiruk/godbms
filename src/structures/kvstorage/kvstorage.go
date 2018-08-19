package kvstorage

import (
	"bytes"
	"encoding/gob"
	"godbms/src/structures"
	"os"
	"time"
)

type KeyValueStorage struct {
	name  string
	cache map[string]KeyValueObject
}

func NewKeyValueStorage(name string, garbageCollectorTick int64) (storage *KeyValueStorage) {
	storage = &KeyValueStorage{name: name}
	storage.load()

	if garbageCollectorTick > 0 {
		storage.garbageCollectorStart(garbageCollectorTick)
	}

	return storage
}

func (storage *KeyValueStorage) garbageCollectorStart(garbageCollectorTick int64) {
	ticker := time.NewTicker(time.Duration(garbageCollectorTick) * time.Millisecond)

	go func() {
		for range ticker.C {
			for k, v := range storage.cache {
				if v.Expire != -1 && v.Expire < time.Now().Unix() {
					delete(storage.cache, k)
				}
			}
		}
	}()
}

func (storage *KeyValueStorage) Get(key string) (value structures.Object) {
	value = storage.cache[key]
	kvObject := value.(KeyValueObject)

	if kvObject.Expire >= 0 && kvObject.Expire < time.Now().Unix() {
		value = KeyValueObject{}
		storage.Delete(key)
	}

	return value
}

func (storage *KeyValueStorage) Set(key string, value string, expire int64) {
	if expire >= 0 {
		expire = time.Now().Unix() + expire
	}

	storage.cache[key] = KeyValueObject{key, value, expire}
	storage.save()
}

func (storage *KeyValueStorage) Delete(key string) {
	delete(storage.cache, key)
}

func (storage *KeyValueStorage) load() {
	f, err := os.Open("./" + storage.name + ".data")

	defer f.Close()

	if err != nil {
		storage.cache = make(map[string]KeyValueObject)
		return
	}

	gob.NewDecoder(f).Decode(&storage.cache)
}

func (storage *KeyValueStorage) save() {
	encodedCache := new(bytes.Buffer)
	gob.NewEncoder(encodedCache).Encode(storage.cache)

	f, _ := os.Create("./" + storage.name + ".data")

	f.Write(encodedCache.Bytes())
	f.Close()
}
