package middleware

import (
	"fmt"
	"net/http"
)

func Logger(debug bool, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if debug {
			fmt.Println(r)
		} else {
			fmt.Println(r.Method + " " + r.URL.Path + " from " + r.RemoteAddr)
		}

		h.ServeHTTP(w, r)
	})
}
