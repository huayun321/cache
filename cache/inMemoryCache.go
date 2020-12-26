package cache

import "sync"

type inMemoryCache struct {
	store map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (i *inMemoryCache) Set(key string, value []byte) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	k, exist := i.store[key]
	if exist {
		i.del(key, k)
	}
	i.store[key] = value
	i.add(key, value)

	return nil
}

func (i *inMemoryCache) Get(key string) ([]byte, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.store[key], nil
}

func (i *inMemoryCache) Del(key string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	k, exist := i.store[key]
	if exist {
		delete(i.store, key)
		i.del(key, k)
	}

	return nil
}
func (i *inMemoryCache) GetStat() Stat {
	return i.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
