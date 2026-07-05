package storage

import "sync"

type Storage struct {
	mtx  sync.RWMutex
	data map[string]string
}

func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Set(key, value string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.data[key] = value
}

func (s *Storage) Get(key string) (string, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	value, ok := s.data[key]
	return value, ok
}

func (s *Storage) Delete(key string) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.data[key]; !ok {
		return false
	}

	delete(s.data, key)
	return true
}

func (s *Storage) Exists(key string) bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	_, ok := s.data[key]
	return ok
}
