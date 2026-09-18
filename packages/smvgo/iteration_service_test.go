package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterationServiceFuncNext(t *testing.T) {
	// Success (1 condition met)
	iS := &iterationService{
		structsMap: make(map[string]struct{}),
	}
	uuid := iS.Next()
	secondUuid := iS.Next()
	assert.True(t, uuid != secondUuid)
}

func TestIterationServiceFuncInstanceIterationService(t *testing.T) {
	// Success
	iIS := InstanceIterationService().(*iterationService)
	secondIIS := InstanceIterationService().(*iterationService)
	iIS.structsMap["key1"] = struct{}{}
	secondIIS.structsMap["key2"] = struct{}{}
	assert.Equal(t, map[string]struct{}{"key1": struct{}{}, "key2": struct{}{}}, iIS.structsMap)
	assert.Equal(t, map[string]struct{}{"key1": struct{}{}, "key2": struct{}{}}, secondIIS.structsMap)
}
