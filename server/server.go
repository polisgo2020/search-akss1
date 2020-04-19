package server

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"searchera/index"
	"time"
)

func Run(addr string, idx index.ReverseIdx, idxPath string) error {
	rootTpl, err := template.ParseFiles("server/root.html")
	if err != nil {
		return err
	}

	resultTpl, err := template.ParseFiles("server/result.html")
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, rootTpl)
	}).Methods("GET")

	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		SearchHandler(w, r, idx, idxPath, resultTpl)
	}).Methods("GET")

	logger := logMiddleware(r)

	srv := &http.Server{
		Handler:      logger,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()

	return err
}

// logMiddleware logging all request
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Debug().
			Str("method", r.Method).
			Str("remote", r.RemoteAddr).
			Str("path", r.URL.Path).
			Int("duration", int(time.Since(start))).
			Msgf("Called url %s", r.URL.Path)
	})
}
