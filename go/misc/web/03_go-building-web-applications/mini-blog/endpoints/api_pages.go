package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/db"
	"net/http"
)

func APIPagesGET(w http.ResponseWriter, r *http.Request) {
	page, err := db.GetPagesData()
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}

	output, err := json.Marshal(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(output))
}
