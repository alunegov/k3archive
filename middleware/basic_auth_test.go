package middleware

import (
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {
	var handleStub httprouter.Handle

	res := BasicAuth(handleStub)

	assert.NotNil(t, res)

	// TODO: res(nil, nil, nil)
}
