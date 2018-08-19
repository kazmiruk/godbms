package kvstorage

type KeyValueObject struct {
	Key    string
	Value  string
	Expire int64
}
