package server

import (
	"fmt"
	"log"
	"net/http"
	"searchera/index"
)

func SearchHandler(w http.ResponseWriter, r *http.Request, idx index.ReverseIdx, idxPath string) {
	log.Println("SearchHandler: ", r.URL.Query())

	resp := ""

	str := r.URL.Query().Get("str")

	resp = fmt.Sprintf("Str: %s\n", str)
	resp += fmt.Sprintf("Idx: %s\n\n", idxPath)

	if str == "" {
		resp += fmt.Sprintf("Str: Invalid parameter")

		_, err := fmt.Fprint(w, resp)
		if err != nil {
			log.Println("Error sending response: " + err.Error())
			return
		}
		return
	}

	found := idx.Search(str)
	if len(found) == 0 {
		resp += fmt.Sprintf("Input string's tokens not found in reverse index")

		_, err := fmt.Fprint(w, resp)
		if err != nil {
			log.Println("Error sending response: " + err.Error())
			return
		}
		return
	}

	result := ""
	for k, v := range found {
		result += fmt.Sprintf("\t%s - %d matches\n", k, v)
	}

	resp += fmt.Sprintf("Results: \n%s\n", result)

	_, err := fmt.Fprint(w, resp)
	if err != nil {
		log.Println("Error sending response: " + err.Error())
		return
	}
}
