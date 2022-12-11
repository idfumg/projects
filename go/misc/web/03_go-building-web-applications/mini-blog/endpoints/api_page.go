package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/db"
	"myapp/session"
	"net/http"

	"github.com/gorilla/mux"
)

func APIPageGET(w http.ResponseWriter, r *http.Request) {
	pageGUID, ok := mux.Vars(r)["guid"]
	if !ok {
		http.Error(w, "Wrong GUID value", http.StatusBadRequest)
		return
	}

	page, err := db.GetPageData(pageGUID, session.UserSession)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	comments, err := db.GetCommentsData(page.Id)
	if err != nil {
		log.Println(err)
	}

	page.Comments = comments

	output, err := json.Marshal(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(output))
}
