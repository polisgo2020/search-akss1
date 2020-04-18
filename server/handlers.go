package server

import (
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"searchera/index"
)

func RootHandler(w http.ResponseWriter, r *http.Request, tpl *template.Template) {
	if err := tpl.Execute(w, nil); err != nil {
		log.Error().Err(err)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request, idx index.ReverseIdx, idxPath string, tpl *template.Template) {
	q := r.URL.Query().Get("q")

	found := idx.Search(q)

	err := tpl.Execute(w, struct {
		Results index.Freq
		Query   string
	}{
		Results: found,
		Query:   q,
	})
	if err != nil {
		log.Error().Err(err)
	}
}
