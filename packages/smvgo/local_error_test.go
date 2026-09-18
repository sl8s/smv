package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalErrorFuncError(t *testing.T) {
	// Success
	lE := localError{
		source:     "LocalErrorTest",
		enumGuilty: enumGuilty{value: "Developer"},
		message:    "Ops",
	}
	assert.Equal(
		t,
		"LocalError{source: LocalErrorTest, enumGuilty: Developer, message: Ops}",
		lE.Error(),
	)
}

func TestLocalErrorFuncEnumGuilty(t *testing.T) {
	// Success
	lE := localError{
		source:     "LocalErrorTest",
		enumGuilty: enumGuilty{value: "Developer"},
		message:    "Ops",
	}
	eG := lE.EnumGuilty().(enumGuilty)
	assert.Equal(
		t,
		"Developer",
		eG.value,
	)
}

func TestLocalErrorFuncMessage(t *testing.T) {
	// Success
	lE := localError{
		source:     "LocalErrorTest",
		enumGuilty: enumGuilty{value: "Developer"},
		message:    "Ops",
	}
	assert.Equal(
		t,
		"Ops",
		lE.Message(),
	)
}

func TestLocalErrorFuncNewLocalError(t *testing.T) {
	// Success
	nLE := NewLocalError(
		"LocalErrorTest",
		enumGuilty{value: "Developer"},
		"Ops",
	).(localError)
	eG := nLE.enumGuilty.(enumGuilty)
	assert.Equal(
		t,
		"LocalErrorTest",
		nLE.source,
	)
	assert.Equal(
		t,
		"Developer",
		eG.value,
	)
	assert.Equal(
		t,
		"Ops",
		nLE.message,
	)
}
