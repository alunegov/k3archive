package middleware

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "github.com/alunegov/k3archive/mocks/net/http"
)

func TestLogger(t *testing.T) {
	handlerMock := &mocks.Handler{}
	handlerMock.On("ServeHTTP", mock.Anything, mock.Anything)

	// debug
	res := Logger(true, handlerMock)
	assert.NotNil(t, res)
	res.ServeHTTP(nil, nil)
	// TODO: examine stdout (or move to log) - expect "<nil>"

	// not debug
	res = Logger(false, handlerMock)
	assert.NotNil(t, res)
	res.ServeHTTP(nil, &http.Request{
		Method:     "method",
		URL:        &url.URL{Path: "url.path"},
		RemoteAddr: "remoteAddr",
	})
	// TODO: examine stdout (or move to log) - expect "method url.path from remoteAddr"
}
