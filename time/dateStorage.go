package time

import (
	"github.com/kerwindena/overtime"

	"sync"
)

type DateStorage struct {
	storage map[int]*Date
	rwMutex sync.RWMutex
}

func NewDateStorage() *DateStorage {
	return &DateStorage{
		storage: make(map[int]*Date),
	}
}

func (ds *DateStorage) Date(day int) overtime.Date {
	ds.rwMutex.RLock()
	date, ok := ds.storage[day]
	ds.rwMutex.RUnlock()

	if ok {
		return date
	}

	ds.rwMutex.Lock()
	defer ds.rwMutex.Unlock()

	date = newDate(day)
	ds.storage[day] = date
	return date
}
