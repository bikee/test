package kvs

import (
	"sync"

	"handouts/lab2/cmd"
)

type KeyValueStorage struct {
	Storage map[string]string
	sync.RWMutex
}

func NewKvs() *KeyValueStorage {
	return &KeyValueStorage{
		Storage: make(map[string]string),
	}
}

func (kvs *KeyValueStorage) Read(key string, resp *cmd.Response) error {
	return nil
}

func (kvs *KeyValueStorage) Write(req *cmd.Request, resp *cmd.Response) error {
	return nil
}
