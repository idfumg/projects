package http

import (
	"myapp/internal/comment"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error retrieving all comments", err)
		return
	}

	if err = sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := (&comment.Comment{}).FromJson(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error retrieving from Json", err)
		return
	}

	comment, err = h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Error posting new commit", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse uint from id", err)
		return
	}

	err = h.Service.DeleteComment(int(id))
	if err != nil {
		sendErrorResponse(w, "Error deleting commit", err)
		return
	}

	if err := sendOkResponse(w, Response{Message: "Successfully deleted comment"}); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse uint from id", err)
		return
	}

	c, err := (&comment.Comment{}).FromJson(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error retrieving from Json", err)
		return
	}

	comment, err := h.Service.UpdateComment(int(id), c)
	if err != nil {
		sendErrorResponse(w, "Error updating comment", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse uint from id", err)
		return
	}

	comment, err := h.Service.GetComment(int(id))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by id", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}
