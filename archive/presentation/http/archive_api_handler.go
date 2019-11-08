package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/alunegov/k3archive/archive"
	"github.com/alunegov/k3archive/middleware"
)

type archiveApiHandler struct {
	usecase archive.Usecase
}

func UseArchiveApiHandler(mux *httprouter.Router, uc archive.Usecase, path string) {
	handler := &archiveApiHandler{
		usecase: uc,
	}

	mux.GET("/api/"+path, handler.getDir)

	mux.GET("/api/"+path+"/:id", handler.getFile)
	mux.DELETE("/api/"+path+"/:id", middleware.BasicAuth(handler.deleteFile))
}

func (it *archiveApiHandler) getDir(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filesInfo, err := it.usecase.GetDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(filesInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (it *archiveApiHandler) getFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fileInfo, err := it.usecase.GetFile(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(fileInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (it *archiveApiHandler) deleteFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := it.usecase.DeleteFile(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
