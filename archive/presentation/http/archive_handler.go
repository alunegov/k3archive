package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UseArchiveHandler(mux *httprouter.Router, fsRoot string, path string) {
	mux.ServeFiles("/"+path+"/*filepath", http.Dir(fsRoot+"/"+path))
}
