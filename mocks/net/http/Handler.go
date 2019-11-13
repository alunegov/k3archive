package mocks

import (
	"net/http"

	mock "github.com/stretchr/testify/mock"
)

type Handler struct {
	mock.Mock
}

func (it *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	it.Called(w, r)
}
