package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/rms/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "c:/Work/Korsar3/Korsar3RPi/Debug/Debug/data/rms/1")
	})

	http.HandleFunc("/t2", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "errr", http.StatusInternalServerError)
	})

	http.ListenAndServe("localhost:3100", nil)
}
