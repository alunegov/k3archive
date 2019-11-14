package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	mocks "github.com/alunegov/k3archive/mocks/archive"
)

func TestGetDir(t *testing.T) {
	useCaseMock := &mocks.Usecase{}
	useCaseMock.On("GetDir").Return(nil, nil)

	sut := &archiveApiHandler{usecase: useCaseMock}

	w := httpCall(sut.getDir, nil)

	useCaseMock.AssertCalled(t, "GetDir")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
	assert.NotEqual(t, 0, w.Body.Len())
}

func TestGetDir_ThenGetDirError(t *testing.T) {
	useCaseMock := &mocks.Usecase{}
	useCaseMock.On("GetDir").Return(nil, errors.New("expected error"))

	sut := &archiveApiHandler{usecase: useCaseMock}

	w := httpCall(sut.getDir, nil)

	useCaseMock.AssertCalled(t, "GetDir")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, w.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	assert.NotEqual(t, 0, w.Body.Len())
}

func TestGetFile(t *testing.T) {
	useCaseMock := &mocks.Usecase{}
	useCaseMock.On("GetFile", "99").Return(nil, nil)

	sut := &archiveApiHandler{usecase: useCaseMock}

	w := httpCall(sut.getFile, []httprouter.Param{{Key: "id", Value: "99"}})

	useCaseMock.AssertCalled(t, "GetFile", "99")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
	assert.NotEqual(t, 0, w.Body.Len())
}

func TestGetFile_ThenGetDirError(t *testing.T) {
	useCaseMock := &mocks.Usecase{}
	useCaseMock.On("GetFile", "99").Return(nil, errors.New("expected error"))

	sut := &archiveApiHandler{usecase: useCaseMock}

	w := httpCall(sut.getFile, []httprouter.Param{{Key: "id", Value: "99"}})

	useCaseMock.AssertCalled(t, "GetFile", "99")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, w.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	assert.NotEqual(t, 0, w.Body.Len())
}

func TestDeleteFile(t *testing.T) {
	useCaseMock := &mocks.Usecase{}
	useCaseMock.On("DeleteFile", "99").Return(nil)

	sut := &archiveApiHandler{usecase: useCaseMock}

	w := httpCall(sut.deleteFile, []httprouter.Param{{Key: "id", Value: "99"}})

	useCaseMock.AssertCalled(t, "DeleteFile", "99")
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, 0, w.Body.Len())
}

func TestDeleteFile_ThenDeleteFileError(t *testing.T) {
	useCaseMock := &mocks.Usecase{}
	useCaseMock.On("DeleteFile", "99").Return(errors.New("expected error"))

	sut := &archiveApiHandler{usecase: useCaseMock}

	w := httpCall(sut.deleteFile, []httprouter.Param{{Key: "id", Value: "99"}})

	useCaseMock.AssertCalled(t, "DeleteFile", "99")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, w.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	assert.NotEqual(t, 0, w.Body.Len())
}

func httpCall(handle httprouter.Handle, params httprouter.Params) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	handle(w, nil, params)
	return w
}
