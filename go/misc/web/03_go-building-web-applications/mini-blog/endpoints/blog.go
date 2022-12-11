package endpoints

import (
	"io"
	"log"
	"myapp/data"
	"myapp/db"
	"myapp/session"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func SendBlogPage(w io.Writer, page *data.Page) error {
	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		return err
	}
	t.Execute(w, page)
	return nil
}

func ServeBlogPage(w http.ResponseWriter, r *http.Request) {
	session.ValidateSession(w, r)

	pageGUID, ok := mux.Vars(r)["guid"]
	if !ok {
		http.Error(w, "Wrong GUID value", http.StatusBadRequest)
		return
	}

	page, err := db.GetPageData(pageGUID, session.UserSession)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}

	err = SendBlogPage(w, page)
	if err != nil {
		http.Error(w, "Couldn't find blog.html", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
