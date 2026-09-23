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
	mutex        sync.Mutex
	anysMap      map[string]any
	listenersMap map[string]map[string]func(event any)
}

var instanceShareService = &shareService{
	anysMap:      make(map[string]any),
	listenersMap: make(map[string]map[string]func(event any)),
}

func (ss *shareService) GetValue(key string, defaultValue any) any {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	value, existsValue := ss.anysMap[key]
	if !existsValue {
		return defaultValue
	}
	return value
}

func (ss *shareService) AddListener(key string, listenerId string, callback func(event any)) error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	listeners, existsListeners := ss.listenersMap[key]
	if !existsListeners {
		ss.listenersMap[key] = make(map[string]func(event any))
		ss.listenersMap[key][listenerId] = callback
		return nil
	}
	_, existsListener := listeners[listenerId]
	if existsListener {
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
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	listeners, existsListeners := ss.listenersMap[key]
	if !existsListeners {
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key to notify listeners: %s", key),
		)
	}
	listener, existsListener := listeners[listenerId]
	if !existsListener {
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key and listenerId to notify listeners: %s -- %s", key, listenerId),
		)
	}
	listener(value)
	return nil
}

func (ss *shareService) NotifyListeners(key string, value any) error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	listeners, existsListeners := ss.listenersMap[key]
	if !existsListeners {
		return NewLocalError(
			"ShareService",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("No exists key to notify listeners: %s", key),
		)
	}
	for _, listener := range listeners {
		listener(value)
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
