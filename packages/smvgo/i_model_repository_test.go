package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type dog struct {
	id   string
	name string
}

func (d dog) Id() string {
	return d.id
}

func (d dog) Clone() IModel {
	return dog{
		id:   d.id,
		name: d.name,
	}
}

func (d dog) ToMap() map[string]any {
	return map[string]any{"id": d.id, "name": d.name}
}

type dogs struct {
	array []dog
}

func (ds *dogs) Models() []dog {
	return ds.array
}

func (ds *dogs) SetModels(array []dog) {
	ds.array = array
}

func (ds *dogs) Clone() IModels[dog] {
	array := ds.array
	newArray := make([]dog, 0, len(array))
	for _, item := range array {
		newArray = append(newArray, item.Clone().(dog))
	}
	return &dogs{array: newArray}
}

func (ds *dogs) ToMaps() []map[string]any {
	array := ds.array
	mapsArray := make([]map[string]any, 0, len(array))
	for _, item := range array {
		mapsArray = append(mapsArray, item.ToMap())
	}
	return mapsArray
}

type dogRepository struct {
	hasClose bool
}

func (dR *dogRepository) Dispose() {
	dR.hasClose = true
}

func (dR *dogRepository) FromMap(m map[string]any) dog {
	return dog{
		id:   m["id"].(string),
		name: m["name"].(string),
	}
}

func (dR *dogRepository) FromMaps(ms []map[string]any) dogs {
	array := make([]dog, 0, len(ms))
	for _, m := range ms {
		array = append(array, dR.FromMap(m))
	}
	return dogs{array: array}
}

func TestIModelRepositoryFuncDispose(t *testing.T) {
	// Success
	dR := &dogRepository{hasClose: false}
	dR.Dispose()
	assert.True(t, dR.hasClose)
}

func TestIModelRepositoryFuncFromMap(t *testing.T) {
	// Success
	dR := &dogRepository{hasClose: false}
	d := dR.FromMap(map[string]any{"id": "id1", "name": "name1"})
	assert.Equal(t, "id1", d.id)
	assert.Equal(t, "name1", d.name)
}

func TestIModelRepositoryFuncFromMaps(t *testing.T) {
	// Success
	dR := &dogRepository{hasClose: false}
	ds := dR.FromMaps([]map[string]any{
		{
			"id":   "id0",
			"name": "name0",
		},
		{
			"id":   "id1",
			"name": "name1",
		},
	})
	assert.Equal(
		t,
		[]dog{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		ds.array,
	)
}

func TestIModelRepositoryFuncGetSafeValue(t *testing.T) {
	// Success
	assert.Equal(
		t,
		"id1",
		GetSafeValue[string](map[string]any{"id": "id1"}, "id", "default"),
	)
	// First condition
	assert.Equal(
		t,
		"default",
		GetSafeValue[string](map[string]any{"id": "id1"}, "name", "default"),
	)
	// Second condition
	assert.Equal(
		t,
		"default",
		GetSafeValue[string](map[string]any{"id": 200}, "id", "default"),
	)
}
