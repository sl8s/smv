package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type product struct {
	id    string
	price int
}

func (p product) Id() string {
	return p.id
}

func (p product) Clone() IModel {
	return product{
		id:    p.id,
		price: p.price,
	}
}

func (p product) ToMap() map[string]any {
	return map[string]any{"id": p.id, "price": p.price}
}

func TestIModelFuncId(t *testing.T) {
	// Success
	p := product{"id1", 100}
	assert.Equal(t, "id1", p.Id())
}

func TestIModelFuncClone(t *testing.T) {
	// Success
	p := product{"id1", 100}
	clonedP := p.Clone().(product)
	assert.Equal(t, "id1", p.id)
	assert.Equal(t, 100, p.price)
	assert.Equal(t, "id1", clonedP.id)
	assert.Equal(t, 100, clonedP.price)
}

func TestIModelFuncToMap(t *testing.T) {
	// Success
	p := product{"id1", 100}
	assert.Equal(t, map[string]any{"id": "id1", "price": 100}, p.ToMap())
}
