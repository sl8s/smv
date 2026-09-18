package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShareServiceFuncGetValue(t *testing.T) {
	// Success
	sS := &shareService{
		anysMap:      map[string]any{"key1": "value1"},
		listenersMap: make(map[string]map[string]func(event any)),
	}
	assert.Equal(t, "value1", sS.GetValue("key1", "default"))
	// First condition
	secondSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: make(map[string]map[string]func(event any)),
	}
	assert.Equal(t, "default", secondSS.GetValue("key1", "default"))
}

func TestShareServiceFuncAddListener(t *testing.T) {
	// Success
	sS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": make(map[string]func(event any))},
	}
	error := sS.AddListener("key1", "listenerId1", func(event any) {})
	_, existsListener := sS.listenersMap["key1"]["listenerId1"]
	assert.True(t, error == nil)
	assert.True(t, existsListener)
	// First condition
	secondSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: make(map[string]map[string]func(event any)),
	}
	secondError := secondSS.AddListener("key1", "listenerId1", func(event any) {})
	_, secondExistsListener := secondSS.listenersMap["key1"]["listenerId1"]
	assert.True(t, secondError == nil)
	assert.True(t, secondExistsListener)
	// Second condition
	thirdSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}}},
	}
	thirdError := thirdSS.AddListener("key1", "listenerId1", func(event any) {})
	assert.True(t, thirdError != nil)
}

func TestShareServiceFuncDeleteListener(t *testing.T) {
	// Success
	sS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}}},
	}
	error := sS.DeleteListener("key1", "listenerId1")
	_, existsListeners := sS.listenersMap["key1"]
	assert.True(t, error == nil)
	assert.False(t, existsListeners)
	// First condition
	secondSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: make(map[string]map[string]func(event any)),
	}
	secondError := secondSS.DeleteListener("key1", "listenerId1")
	assert.True(t, secondError != nil)
	// Second condition
	thirdSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}, "listenerId2": func(event any) {}}},
	}
	thirdError := thirdSS.DeleteListener("key1", "listenerId1")
	_, thirdExistsListeners := thirdSS.listenersMap["key1"]
	assert.True(t, thirdError == nil)
	assert.True(t, thirdExistsListeners)
}

func TestShareServiceFuncNotifyListener(t *testing.T) {
	// Success
	eventFromListener := ""
	listener := func(event any) {
		eventFromListener = event.(string)
	}
	sS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": listener}},
	}
	error := sS.NotifyListener("key1", "listenerId1", "test")
	assert.True(t, error == nil)
	assert.Equal(t, "test", eventFromListener)
	// First condition
	secondSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: make(map[string]map[string]func(event any)),
	}
	secondError := secondSS.NotifyListener("key1", "listenerId1", "test")
	assert.True(t, secondError != nil)
	// Second condition
	thirdSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": make(map[string]func(event any))},
	}
	thirdError := thirdSS.NotifyListener("key1", "listenerId1", "test")
	assert.True(t, thirdError != nil)
}

func TestShareServiceFuncNotifyListeners(t *testing.T) {
	// Success
	eventFromFirstListener := ""
	firstListener := func(event any) {
		eventFromFirstListener = event.(string)
	}
	eventFromSecondListener := ""
	secondListener := func(event any) {
		eventFromSecondListener = event.(string)
	}
	sS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": firstListener, "listenerId2": secondListener}},
	}
	error := sS.NotifyListeners("key1", "test")
	assert.True(t, error == nil)
	assert.Equal(t, "test", eventFromFirstListener)
	assert.Equal(t, "test", eventFromSecondListener)
	// First condition
	secondSS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: make(map[string]map[string]func(event any)),
	}
	secondError := secondSS.NotifyListeners("key1", "test")
	assert.True(t, secondError != nil)
}

func TestShareServiceFuncDeleteListeners(t *testing.T) {
	// Success
	sS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}, "listenerId2": func(event any) {}}},
	}
	sS.DeleteListeners("key1")
	assert.Equal(t, map[string]map[string]func(event any){}, sS.listenersMap)
}

func TestShareServiceFuncUpdate(t *testing.T) {
	// Success
	sS := &shareService{
		anysMap:      make(map[string]any),
		listenersMap: make(map[string]map[string]func(event any)),
	}
	sS.Update("key1", "value1")
	assert.Equal(t, "value1", sS.anysMap["key1"])
}

func TestShareServiceFuncDelete(t *testing.T) {
	// Success
	sS := &shareService{
		anysMap:      map[string]any{"key1": "value1"},
		listenersMap: make(map[string]map[string]func(event any)),
	}
	sS.Delete("key1")
	_, existsAny := sS.anysMap["key1"]
	assert.False(t, existsAny)
}

func TestShareServiceFuncInstanceShareService(t *testing.T) {
	// Success
	iSS := InstanceShareService().(*shareService)
	secondISS := InstanceShareService().(*shareService)
	iSS.anysMap["key1"] = "value1"
	secondISS.anysMap["key2"] = "value2"
	assert.Equal(t, map[string]any{"key1": "value1", "key2": "value2"}, iSS.anysMap)
	assert.Equal(t, map[string]any{"key1": "value1", "key2": "value2"}, secondISS.anysMap)
}
