// Go Workshop - Practise - Gube - Storage

package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/themue/workshop-go/practise/gube/pkg/actor"
)

// InMemoryStorage implements the Storage interface using a
// simple in-memory map. It provides concurrent usage.
type InMemoryStorage struct {
	data      map[string]string
	listeners map[string]Listener
	act       *actor.Actor
}

// NewInMemoryStorage creates an instance of the in-memory
// storage.
func NewInMemoryStorage(ctx context.Context) *InMemoryStorage {
	return &InMemoryStorage{
		data:      make(map[string]string),
		listeners: make(map[string]Listener),
		act:       actor.New(actor.Context(ctx)),
	}
}

// Put implements the Storage interface.
func (s *InMemoryStorage) Put(key, value string) (string, error) {
	var out string

	err := s.act.DoSync(func() {
		out = s.data[key]
		s.data[key] = value
		for _, listen := range s.listeners {
			go listen(PutOp, key, value)
		}
	}, time.Second)

	return out, err
}

// Get implements the Storage interface.
func (s *InMemoryStorage) Get(key string) (string, error) {
	var out string
	var gerr error

	if err := s.act.DoSync(func() {
		var ok bool
		out, ok = s.data[key]
		if !ok {
			gerr = fmt.Errorf("no value found for key %q", key)
		}
	}, time.Second); err != nil {
		return out, err
	}

	return out, gerr
}

// Delete implements the Storage interface.
func (s *InMemoryStorage) Delete(key string) (string, error) {
	var out string
	var gerr error

	if err := s.act.DoSync(func() {
		var ok bool
		out, ok = s.data[key]
		if !ok {
			gerr = fmt.Errorf("no value found for key %q", key)
			return
		}
		delete(s.data, key)
		for _, listen := range s.listeners {
			go listen(DeleteOp, key, out)
		}
	}, time.Second); err != nil {
		return out, err
	}

	return out, gerr
}

// Find implements the Registry interface.
func (s *InMemoryStorage) Find(matches func(key, value string) bool) (map[string]string, error) {
	var outs = map[string]string{}
	var ferr error

	if err := s.act.DoSync(func() {
		for key, value := range s.data {
			if matches(key, value) {
				outs[key] = value
			}
		}
	}, 2*time.Second); err != nil {
		return outs, err
	}

	return outs, ferr
}

// AddListener implements the Storage interface.
func (s *InMemoryStorage) AddListener(id string, l Listener) {
	if err := s.act.DoAsync(func() {
		s.listeners[id] = l
	}, 30*time.Second); err != nil {
		// Hard failure!
		panic(err)
	}
}

// RemoveListener implements the Storage interface.
func (s *InMemoryStorage) RemoveListener(id string) {
	if err := s.act.DoAsync(func() {
		delete(s.listeners, id)
	}, 30*time.Second); err != nil {
		// Hard failure!
		panic(err)
	}
}
