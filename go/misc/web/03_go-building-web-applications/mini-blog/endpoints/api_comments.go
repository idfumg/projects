package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddCommentResponse struct {
	Fields map[string]string
}

type UpdateCommentResponse struct {
	Fields map[string]string
}

func NewAddCommentResponse(id int64, wasAdded bool) *AddCommentResponse {
	return &AddCommentResponse{
		Fields: map[string]string{
			"id":        fmt.Sprint(id),
			"was_added": strconv.FormatBool(wasAdded),
		},
	}
}

func NewUpdateCommentResponse(id string, wasUpdated bool) *UpdateCommentResponse {
	return &UpdateCommentResponse{
		Fields: map[string]string{
			"id":          id,
			"was_updated": strconv.FormatBool(wasUpdated),
		},
	}
}

func APIAddCommentPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	pageId := r.FormValue("pageId")
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")

	id, wasAdded, err := db.CreateNewCommentData(pageId, name, email, comments)
	if err != nil {
		log.Println(err)
		return
	}

	resp := NewAddCommentResponse(id, wasAdded)
	output, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(output))
}

func APIUpdateCommentPUT(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	id, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "Wrong id value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")

	_, wasUpdated, err := db.UpdateCommentData(id, name, email, comments)
	if err != nil {
		log.Println(err)
		return
	}

	resp := NewUpdateCommentResponse(id, wasUpdated)
	output, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(output))
}
