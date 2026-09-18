package smvgo

import (
	"sync"
	"uuid"
)

type IterationService interface {
	Next() string
	final()
}

type iterationService struct {
	mutex      sync.RWMutex
	structsMap map[string]struct{}
}

var instanceIterationService = &iterationService{
	structsMap: make(map[string]struct{}),
}

func (is *iterationService) Next() string {
	is.mutex.Lock()
	defer is.mutex.Unlock()
	uuidVersionFour := ""
	for {
		uuidVersionFour = uuid.NewV4().String()
		if _, existsStruct := is.structsMap[uuidVersionFour]; !existsStruct {
			break
		}
	}
	is.structsMap[uuidVersionFour] = struct{}{}
	return uuidVersionFour
}

func (is *iterationService) final() {
}

func InstanceIterationService() IterationService {
	return instanceIterationService
}
