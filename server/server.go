package server

import (
	"github.com/gorilla/mux"
	"html/template"
	"net"
	"net/http"
	"searchera/index"
	"time"
)

func Run(host string, port string, idx index.ReverseIdx, idxPath string) error {
	r := mux.NewRouter()

	rootTpl, err := template.ParseFiles("server/root.html")
	if err != nil {
		return err
	}

	resultTpl, err := template.ParseFiles("server/result.html")
	if err != nil {
		return err
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, rootTpl)
	}).Methods("GET")

	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		SearchHandler(w, r, idx, idxPath, resultTpl)
	}).Methods("GET")

	addr := net.JoinHostPort(host, port)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()

	return err
}
