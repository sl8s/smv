package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type cat struct {
	id   string
	name string
}

func (c cat) Id() string {
	return c.id
}

func (c cat) Clone() IModel {
	return cat{
		id:   c.id,
		name: c.name,
	}
}

func (c cat) ToMap() map[string]any {
	return map[string]any{"id": c.id, "name": c.name}
}

type cats struct {
	array []cat
}

func (cs *cats) Models() []cat {
	return cs.array
}

func (cs *cats) SetModels(array []cat) {
	cs.array = array
}

func (cs *cats) Clone() IModels[cat] {
	array := cs.array
	newArray := make([]cat, 0, len(array))
	for _, item := range array {
		newArray = append(newArray, item.Clone().(cat))
	}
	return &cats{array: newArray}
}

func (cs *cats) ToMaps() []map[string]any {
	array := cs.array
	mapsArray := make([]map[string]any, 0, len(array))
	for _, item := range array {
		mapsArray = append(mapsArray, item.ToMap())
	}
	return mapsArray
}

func TestIModelsFuncModels(t *testing.T) {
	// Success
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
		},
	}
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
		},
		cs.Models(),
	)
}

func TestIModelsFuncSetModels(t *testing.T) {
	// Success
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
		},
	}
	cs.SetModels([]cat{
		{id: "id5", name: "name5"},
		{id: "id6", name: "name6"},
		{id: "id7", name: "name7"},
		{id: "id8", name: "name8"},
		{id: "id9", name: "name9"},
	},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id5", name: "name5"},
			{id: "id6", name: "name6"},
			{id: "id7", name: "name7"},
			{id: "id8", name: "name8"},
			{id: "id9", name: "name9"},
		},
		cs.array,
	)
}

func TestIModelsFuncClone(t *testing.T) {
	// Success
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
		},
	}
	clonedCs := cs.Clone().(*cats)
	cs.array = append(cs.array, cat{id: "id5", name: "name5"})
	clonedCs.array[2] = cat{id: "id2", name: "updatedName2"}
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
			{id: "id5", name: "name5"},
		},
		cs.array,
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "updatedName2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
		},
		clonedCs.array,
	)
}

func TestIModelsFuncToMaps(t *testing.T) {
	// Success
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
			{id: "id4", name: "name4"},
		},
	}
	assert.Equal(
		t,
		[]map[string]any{
			{
				"id":   "id0",
				"name": "name0",
			},
			{
				"id":   "id1",
				"name": "name1",
			},
			{
				"id":   "id2",
				"name": "name2",
			},
			{
				"id":   "id3",
				"name": "name3",
			},
			{
				"id":   "id4",
				"name": "name4",
			},
		},
		cs.ToMaps(),
	)
}

func TestIModelsFuncAddFromIModelsAndNewModel(t *testing.T) {
	// Success
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
		},
	}
	error := AddFromIModelsAndNewModel[cat](
		cs,
		cat{id: "id1", name: "name1"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		cs.array,
	)
	assert.True(
		t,
		error == nil,
	)
	// First condition
	secondCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
		},
	}
	secondError := AddFromIModelsAndNewModel[cat](
		secondCs,
		cat{id: "id0", name: "name0"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
		},
		secondCs.array,
	)
	assert.True(
		t,
		secondError != nil,
	)
}

func TestIModelsFuncUpdateFromIModelsAndNewModelById(t *testing.T) {
	// Success (1 condition met)
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
		},
	}
	error := UpdateFromIModelsAndNewModelById[cat](
		cs,
		cat{id: "id0", name: "updatedName0"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "updatedName0"},
			{id: "id0", name: "updatedName0"},
		},
		cs.array,
	)
	assert.True(
		t,
		error == nil,
	)
	// First condition
	secondCs := &cats{array: []cat{}}
	secondError := UpdateFromIModelsAndNewModelById[cat](
		secondCs,
		cat{id: "id", name: "name"},
	)
	assert.Equal(
		t,
		[]cat{},
		secondCs.array,
	)
	assert.True(
		t,
		secondError != nil,
	)
	// Second condition
	thirdCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	thirdError := UpdateFromIModelsAndNewModelById[cat](
		thirdCs,
		cat{id: "id1", name: "updatedName1"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "updatedName1"},
		},
		thirdCs.array,
	)
	assert.True(
		t,
		thirdError == nil,
	)
	// Third condition
	fourthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	fourthError := UpdateFromIModelsAndNewModelById[cat](
		fourthCs,
		cat{id: "id", name: "name"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		fourthCs.array,
	)
	assert.True(
		t,
		fourthError != nil,
	)
}

func TestIModelsFuncDeleteFromIModelsAndIdById(t *testing.T) {
	// Success (1 condition met)
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	error := DeleteFromIModelsAndIdById[cat](
		cs,
		"id0",
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id1", name: "name1"},
		},
		cs.array,
	)
	assert.True(
		t,
		error == nil,
	)
	// First condition
	secondCs := &cats{
		array: []cat{},
	}
	secondError := DeleteFromIModelsAndIdById[cat](
		secondCs,
		"id0",
	)
	assert.Equal(
		t,
		[]cat{},
		secondCs.array,
	)
	assert.True(
		t,
		secondError != nil,
	)
	// Second condition
	thirdCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	thirdError := DeleteFromIModelsAndIdById[cat](
		thirdCs,
		"id",
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		thirdCs.array,
	)
	assert.True(
		t,
		thirdError != nil,
	)
}

func TestIModelsFuncAddFromIModelsAndNewModels(t *testing.T) {
	// Success
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
		},
	}
	error := AddFromIModelsAndNewModels[cat](
		cs,
		[]cat{
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
		},
		cs.array,
	)
	assert.True(
		t,
		error == nil,
	)
	// First condition
	secondCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	secondError := AddFromIModelsAndNewModels[cat](
		secondCs,
		[]cat{},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		secondCs.array,
	)
	assert.True(
		t,
		secondError != nil,
	)
	// Second condition
	thirdCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	thirdError := AddFromIModelsAndNewModels[cat](
		thirdCs,
		[]cat{
			{id: "id2", name: "name2"},
			{id: "id2", name: "name2"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		thirdCs.array,
	)
	assert.True(
		t,
		thirdError != nil,
	)
	// Third condition
	fourthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	fourthError := AddFromIModelsAndNewModels[cat](
		fourthCs,
		[]cat{
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		fourthCs.array,
	)
	assert.True(
		t,
		fourthError != nil,
	)
}

func TestIModelsFuncUpdateFromIModelsAndNewModelsById(t *testing.T) {
	// Success (1 condition met)
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	error := UpdateFromIModelsAndNewModelsById[cat](
		cs,
		[]cat{
			{id: "id0", name: "updatedName0"},
			{id: "id1", name: "updatedName1"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "updatedName0"},
			{id: "id0", name: "updatedName0"},
			{id: "id1", name: "updatedName1"},
		},
		cs.array,
	)
	assert.True(
		t,
		error == nil,
	)
	// First condition and first sub-condition
	secondCs := &cats{
		array: []cat{},
	}
	secondError := UpdateFromIModelsAndNewModelsById[cat](
		secondCs,
		[]cat{
			{id: "id0", name: "updatedName0"},
			{id: "id1", name: "updatedName1"},
		},
	)
	assert.Equal(
		t,
		[]cat{},
		secondCs.array,
	)
	assert.True(
		t,
		secondError != nil,
	)
	// First condition and second sub-condition
	thirdCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	thirdError := UpdateFromIModelsAndNewModelsById[cat](
		thirdCs,
		[]cat{},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		thirdCs.array,
	)
	assert.True(
		t,
		thirdError != nil,
	)
	// Second condition
	fourthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	fourthError := UpdateFromIModelsAndNewModelsById[cat](
		fourthCs,
		[]cat{
			{id: "id2", name: "name2"},
			{id: "id2", name: "name2"},
			{id: "id3", name: "name3"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		fourthCs.array,
	)
	assert.True(
		t,
		fourthError != nil,
	)
	// Third condition
	fifthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name1"},
		},
	}
	fifthError := UpdateFromIModelsAndNewModelsById[cat](
		fifthCs,
		[]cat{
			{id: "id0", name: "updatedName0"},
			{id: "id2", name: "updatedName2"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "updatedName0"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "updatedName2"},
		},
		fifthCs.array,
	)
	assert.True(
		t,
		fifthError == nil,
	)
	// Fourth condition
	sixthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	sixthError := UpdateFromIModelsAndNewModelsById[cat](
		sixthCs,
		[]cat{
			{id: "id2", name: "updatedName2"},
			{id: "id3", name: "updatedName3"},
		},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		sixthCs.array,
	)
	assert.True(
		t,
		sixthError != nil,
	)
}

func TestIModelsFuncDeleteFromIModelsAndIdsById(t *testing.T) {
	// Success (1 condition met)
	cs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
			{id: "id1", name: "name1"},
			{id: "id2", name: "name2"},
		},
	}
	error := DeleteFromIModelsAndIdsById[cat](
		cs,
		[]string{"id0", "id1"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id2", name: "name2"},
		},
		cs.array,
	)
	assert.True(
		t,
		error == nil,
	)
	// First condition and first sub-condition
	secondCs := &cats{
		array: []cat{},
	}
	secondError := DeleteFromIModelsAndIdsById[cat](
		secondCs,
		[]string{"id0", "id1"},
	)
	assert.Equal(
		t,
		[]cat{},
		secondCs.array,
	)
	assert.True(
		t,
		secondError != nil,
	)
	// First condition and second sub-condition
	thirdCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	thirdError := DeleteFromIModelsAndIdsById[cat](
		thirdCs,
		[]string{},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		thirdCs.array,
	)
	assert.True(
		t,
		thirdError != nil,
	)
	// Second condition
	fourthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	fourthError := DeleteFromIModelsAndIdsById[cat](
		fourthCs,
		[]string{"id0", "id1", "id1"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		fourthCs.array,
	)
	assert.True(
		t,
		fourthError != nil,
	)
	// Third condition
	fifthCs := &cats{
		array: []cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
	}
	fifthError := DeleteFromIModelsAndIdsById[cat](
		fifthCs,
		[]string{"id3", "id4"},
	)
	assert.Equal(
		t,
		[]cat{
			{id: "id0", name: "name0"},
			{id: "id1", name: "name1"},
		},
		fifthCs.array,
	)
	assert.True(
		t,
		fifthError != nil,
	)
}
