package main

//go:generate go-bindata-assetfs -pkg main app/...
import (
	"net/http"

	"github.com/gorilla/mux"
)

func BindataServer(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := root + r.RequestURI
		bytes, err := Asset(id)
		if err != nil {
			w.Write([]byte("file not found"))
		} else {
			w.Write(bytes)
		}
	})
}
func addBindata(router *mux.Router) {
	router.PathPrefix("/local").Handler(BindataServer("app"))
}
