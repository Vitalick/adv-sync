package advsync

import "sync"

//NamedMutex is a named mutex via sync.RWMutex
type NamedMutex struct {
	mapLock     sync.RWMutex
	internalMap map[interface{}]*sync.Mutex
}

//NewNamedMutex create new named mutex
func NewNamedMutex() *NamedMutex {
	return &NamedMutex{
		internalMap: map[interface{}]*sync.Mutex{},
	}
}

//Unlock mutex by name
func (nm *NamedMutex) Unlock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.Mutex{}
		nm.internalMap[slug].Unlock()
		nm.mapLock.Unlock()
		return
	}
	locker.Unlock()
}

//Lock mutex by name
func (nm *NamedMutex) Lock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.Mutex{}
		nm.internalMap[slug].Lock()
		nm.mapLock.Unlock()
		return
	}
	locker.Lock()
}
