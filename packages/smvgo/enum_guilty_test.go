package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumGuiltyFuncString(t *testing.T) {
	// Success
	eG := enumGuilty{value: "Test"}
	assert.Equal(t, "Test", eG.String())
}

func TestEnumGuiltyFuncDeveloperByEnumGuilty(t *testing.T) {
	// Success
	eG := DeveloperByEnumGuilty().(enumGuilty)
	assert.Equal(t, "Developer", eG.value)
}

func TestEnumGuiltyFuncDeviceByEnumGuilty(t *testing.T) {
	// Success
	eG := DeviceByEnumGuilty().(enumGuilty)
	assert.Equal(t, "Device", eG.value)
}

func TestEnumGuiltyFuncUserByEnumGuilty(t *testing.T) {
	// Success
	eG := UserByEnumGuilty().(enumGuilty)
	assert.Equal(t, "User", eG.value)
}
