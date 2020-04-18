package server

import (
	"html/template"
	"log"
	"net/http"
	"searchera/index"
)

func RootHandler(w http.ResponseWriter, r *http.Request, tpl *template.Template) {
	log.Println("RootHandler: ", r.URL.Query())

	if err := tpl.Execute(w, nil); err != nil {
		log.Println("RootHandler: ", err.Error())
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request, idx index.ReverseIdx, idxPath string, tpl *template.Template) {
	log.Println("SearchHandler: ", r.URL.Query())

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
		log.Println("Error sending response: ", err.Error())
	}

}
