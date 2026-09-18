package smvgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkErrorFuncError(t *testing.T) {
	// Success
	nE := networkError{
		source:                "NetworkErrorTest",
		statusCode:            403,
		nameStatusCode:        "403 Forbidden",
		descriptionStatusCode: "The server understood the request, but is refusing to fulfill it.",
	}
	assert.Equal(
		t,
		"NetworkError{source: NetworkErrorTest, statusCode: 403, nameStatusCode: 403 Forbidden, descriptionStatusCode: The server understood the request, but is refusing to fulfill it.}",
		nE.Error(),
	)
}

func TestNetworkErrorFuncStatusCode(t *testing.T) {
	// Success
	nE := networkError{
		source:                "NetworkErrorTest",
		statusCode:            403,
		nameStatusCode:        "403 Forbidden",
		descriptionStatusCode: "The server understood the request, but is refusing to fulfill it.",
	}
	assert.Equal(
		t,
		403,
		nE.StatusCode(),
	)
}

func TestNetworkErrorFuncNameStatusCode(t *testing.T) {
	// Success
	nE := networkError{
		source:                "NetworkErrorTest",
		statusCode:            403,
		nameStatusCode:        "403 Forbidden",
		descriptionStatusCode: "The server understood the request, but is refusing to fulfill it.",
	}
	assert.Equal(
		t,
		"403 Forbidden",
		nE.NameStatusCode(),
	)
}

func TestNetworkErrorFuncDescriptionStatusCode(t *testing.T) {
	// Success
	nE := networkError{
		source:                "NetworkErrorTest",
		statusCode:            403,
		nameStatusCode:        "403 Forbidden",
		descriptionStatusCode: "The server understood the request, but is refusing to fulfill it.",
	}
	assert.Equal(
		t,
		"The server understood the request, but is refusing to fulfill it.",
		nE.DescriptionStatusCode(),
	)
}

func TestNetworkErrorFuncNewNetworkError(t *testing.T) {
	// Success
	nNE := NewNetworkError(
		"NetworkErrorTest",
		403,
		"403 Forbidden",
		"The server understood the request, but is refusing to fulfill it.",
	).(networkError)
	assert.Equal(
		t,
		"NetworkErrorTest",
		nNE.source,
	)
	assert.Equal(
		t,
		403,
		nNE.statusCode,
	)
	assert.Equal(
		t,
		"403 Forbidden",
		nNE.nameStatusCode,
	)
	assert.Equal(
		t,
		"The server understood the request, but is refusing to fulfill it.",
		nNE.descriptionStatusCode,
	)
}

func TestNetworkErrorFuncNewNetworkErrorFromSourceAndStatusCode(t *testing.T) {
	// Success
	statusCodes := []int{1000, 201, 202, 203, 204, 205, 206, 300, 301, 302, 303, 304, 305, 307, 400, 401, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 500, 501, 502, 503, 504, 505}
	for _, statusCode := range statusCodes {
		nNE := NewNetworkErrorFromSourceAndStatusCode("NetworkErrorTest", statusCode).(networkError)
		assert.Equal(
			t,
			"NetworkErrorTest",
			nNE.source,
		)
		assert.Equal(
			t,
			statusCode,
			nNE.statusCode,
		)
	}
}
