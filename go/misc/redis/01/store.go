package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Cacher interface {
	Get(int) (string, bool)
	Set(int, string) error
	Remove(int) error
}

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(cacher Cacher) *Store {
	data := map[int]string{
		1: "Elon Musk is the new owner of Twitter",
		2: "Foo is not bar and bar is not baz",
		3: "Must watch Arcane",
	}
	return &Store{
		data:  data,
		cache: cacher,
	}
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		if err := s.cache.Remove(key); err != nil {
			log.Warn(err)
		}
		log.Info("[Redis cache] The value is found")
		return val, nil
	}
	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}
	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}
	log.Info("[Internal storage] The value is found")
	return val, nil
}
