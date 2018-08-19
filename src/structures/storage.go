package structures

type Storage interface {
	Get(key string) (value Object)
	Set(key string, value string, expire int64)
	Delete(key string)
}
