package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShareProxyFuncGetValue(t *testing.T) {
	// Success
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      map[string]any{"key1": "value1"},
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	assert.Equal(t, "value1", sP.GetValue("key1", "default"))
	// First condition
	secondSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	assert.Equal(t, "default", secondSP.GetValue("key1", "default"))
}

func TestShareProxyFuncAddListener(t *testing.T) {
	// Success
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": make(map[string]func(event any))},
		},
	}
	error := sP.AddListener("key1", func(event any) {})
	sS := sP.shareService.(*shareService)
	_, existsListener := sS.listenersMap["key1"]["listenerId1"]
	assert.True(t, error == nil)
	assert.True(t, existsListener)
	// First condition
	secondSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	secondError := secondSP.AddListener("key1", func(event any) {})
	secondSS := secondSP.shareService.(*shareService)
	_, secondExistsListener := secondSS.listenersMap["key1"]["listenerId1"]
	assert.True(t, secondError == nil)
	assert.True(t, secondExistsListener)
	// Second condition
	thirdSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}}},
		},
	}
	thirdError := thirdSP.AddListener("key1", func(event any) {})
	assert.True(t, thirdError != nil)
}

func TestShareProxyFuncDeleteListener(t *testing.T) {
	// Success
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}}},
		},
	}
	error := sP.DeleteListener("key1")
	sS := sP.shareService.(*shareService)
	_, existsListeners := sS.listenersMap["key1"]
	assert.True(t, error == nil)
	assert.False(t, existsListeners)
	// First condition
	secondSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	secondError := secondSP.DeleteListener("key1")
	assert.True(t, secondError != nil)
	// Second condition
	thirdSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}, "listenerId2": func(event any) {}}},
		},
	}
	thirdError := thirdSP.DeleteListener("key1")
	thirdSS := thirdSP.shareService.(*shareService)
	_, thirdExistsListeners := thirdSS.listenersMap["key1"]
	assert.True(t, thirdError == nil)
	assert.True(t, thirdExistsListeners)
}

func TestShareProxyFuncNotifyListener(t *testing.T) {
	// Success
	eventFromListener := ""
	listener := func(event any) {
		eventFromListener = event.(string)
	}
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": listener}},
		},
	}
	error := sP.NotifyListener("key1", "test")
	assert.True(t, error == nil)
	assert.Equal(t, "test", eventFromListener)
	// First condition
	secondSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	secondError := secondSP.NotifyListener("key1", "test")
	assert.True(t, secondError != nil)
	// Second condition
	thirdSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": make(map[string]func(event any))},
		},
	}
	thirdError := thirdSP.NotifyListener("key1", "test")
	assert.True(t, thirdError != nil)
}

func TestShareProxyFuncNotifyListeners(t *testing.T) {
	// Success
	eventFromFirstListener := ""
	firstListener := func(event any) {
		eventFromFirstListener = event.(string)
	}
	eventFromSecondListener := ""
	secondListener := func(event any) {
		eventFromSecondListener = event.(string)
	}
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": firstListener, "listenerId2": secondListener}},
		},
	}
	error := sP.NotifyListeners("key1", "test")
	assert.True(t, error == nil)
	assert.Equal(t, "test", eventFromFirstListener)
	assert.Equal(t, "test", eventFromSecondListener)
	// First condition
	secondSP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	secondError := secondSP.NotifyListeners("key1", "test")
	assert.True(t, secondError != nil)
}

func TestShareProxyFuncDeleteListeners(t *testing.T) {
	// Success
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: map[string]map[string]func(event any){"key1": map[string]func(event any){"listenerId1": func(event any) {}, "listenerId2": func(event any) {}}},
		},
	}
	sP.DeleteListeners("key1")
	sS := sP.shareService.(*shareService)
	assert.Equal(t, map[string]map[string]func(event any){}, sS.listenersMap)
}

func TestShareProxyFuncUpdate(t *testing.T) {
	// Success
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      make(map[string]any),
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	sP.Update("key1", "value1")
	sS := sP.shareService.(*shareService)
	assert.Equal(t, "value1", sS.anysMap["key1"])
}

func TestShareProxyFuncDelete(t *testing.T) {
	// Success
	sP := &shareProxy{
		listenerId: "listenerId1",
		shareService: &shareService{
			anysMap:      map[string]any{"key1": "value1"},
			listenersMap: make(map[string]map[string]func(event any)),
		},
	}
	sP.Delete("key1")
	sS := sP.shareService.(*shareService)
	_, existsAny := sS.anysMap["key1"]
	assert.False(t, existsAny)
}

func TestShareProxyFuncNewShareProxy(t *testing.T) {
	// Success
	nSP := NewShareProxy().(*shareProxy)
	sS := nSP.shareService.(*shareService)
	assert.True(t, nSP.listenerId != "")
	assert.Equal(t, make(map[string]any), sS.anysMap)
	assert.Equal(t, make(map[string]map[string]func(event any)), sS.listenersMap)
}
