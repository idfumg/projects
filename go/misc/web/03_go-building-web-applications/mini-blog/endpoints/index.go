package endpoints

import (
	"io"
	"log"
	"myapp/data"
	"myapp/db"
	"net/http"
	"text/template"
)

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 301)
}

func SendIndexPage(w io.Writer, pages []*data.Page) error {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return err
	}
	t.Execute(w, pages)
	return nil
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	pages, err := db.GetPagesData()
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}

	err = SendIndexPage(w, pages)
	if err != nil {
		http.Error(w, "Couldn't find index.html", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
