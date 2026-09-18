package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type implementDispose struct {
	hasClose bool
}

func (iD *implementDispose) Dispose() {
	iD.hasClose = true
}

func TestIDisposeFuncDispose(t *testing.T) {
	// Success
	iD := &implementDispose{hasClose: false}
	iD.Dispose()
	assert.True(t, iD.hasClose)
}
