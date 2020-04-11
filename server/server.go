package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"searchera/index"
	"time"
)

func Run(addr string, idx index.ReverseIdx, idxPath string) error {
	r := mux.NewRouter()
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		SearchHandler(w, r, idx, idxPath)
	}).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	return err
}
