package smvgo

import (
	"fmt"
	"sync"
)

type ShareService interface {
	GetValue(key string, defaultValue any) any
	AddListener(key string, listenerId string, callback func(event any)) error
	DeleteListener(key string, listenerId string) error
	NotifyListener(key string, listenerId string, value any) error
	NotifyListeners(key string, value any) error
	DeleteListeners(key string)
	Update(key string, value any)
	Delete(key string)
	final()
}

type shareService struct {
	mutex        sync.RWMutex
	anysMap      map[string]any
	listenersMap map[string]map[string]func(event any)
}

var instanceShareService = &shareService{
	anysMap:      make(map[string]any),
	listenersMap: make(map[string]map[string]func(event any)),
}

func (ss *shareService) GetValue(key string, defaultValue any) any {
	ss.mutex.RLock()
	defer ss.mutex.RUnlock()
	value, existsValue := ss.anysMap[key]
	if !existsValue {
		return defaultValue
	}
	return value
}

func (ss *shareService) AddListener(key string, listenerId string, callback func(event any)) error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	if _, existsListeners := ss.listenersMap[key]; !existsListeners {
		ss.listenersMap[key] = make(map[string]func(event any))
		ss.listenersMap[key][listenerId] = callback
		return nil
	}
	if _, existsListener := ss.listenersMap[key][listenerId]; existsListener {
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("Under such a key and listenerId there already exists a listener: %s -- %s", key, listenerId),
		)
	}
	ss.listenersMap[key][listenerId] = callback
	return nil
}

func (ss *shareService) DeleteListener(key string, listenerId string) error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	listeners, existsListeners := ss.listenersMap[key]
	if !existsListeners {
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key to delete listeners: %s", key),
		)
	}
	delete(listeners, listenerId)
	if len(listeners) > 0 {
		return nil
	}
	delete(ss.listenersMap, key)
	return nil
}

func (ss *shareService) NotifyListener(key string, listenerId string, value any) error {
	ss.mutex.RLock()
	listeners, existsListeners := ss.listenersMap[key]
	if !existsListeners {
		ss.mutex.RUnlock()
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key to notify listeners: %s", key),
		)
	}
	listener, existsListener := listeners[listenerId]
	if !existsListener {
		ss.mutex.RUnlock()
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key and listenerId to notify listeners: %s -- %s", key, listenerId),
		)
	}
	ss.mutex.RUnlock()
	listener(value)
	return nil
}

func (ss *shareService) NotifyListeners(key string, value any) error {
	ss.mutex.RLock()
	listeners, existsListeners := ss.listenersMap[key]
	if !existsListeners {
		ss.mutex.RUnlock()
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key to notify listeners: %s", key),
		)
	}
	snapshots := make([]func(event any), 0, len(listeners))
	for _, listener := range listeners {
		snapshots = append(snapshots, listener)
	}
	ss.mutex.RUnlock()
	for _, snapshot := range snapshots {
		snapshot(value)
	}
	return nil
}

func (ss *shareService) DeleteListeners(key string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	delete(ss.listenersMap, key)
}

func (ss *shareService) Update(key string, value any) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.anysMap[key] = value
}

func (ss *shareService) Delete(key string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	delete(ss.anysMap, key)
}

func (ss *shareService) final() {
}

func InstanceShareService() ShareService {
	return instanceShareService
}
