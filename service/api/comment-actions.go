package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// Delete comment from a photo
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	var commentid uint64
	var userid uint64

	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentid, err = strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.PhotoId = photoid
	comment.Id = commentid

	userid, err = rt.db.ObtainCommentUserId(commentid)
	if errors.Is(err, database.ErrCommentDoesNotExist) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthentication(r.Header.Get("Authorization"), userid)
	if errors.Is(err, database.ErrNotAuthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbcomment := comment.CommentFromApiToDatabase()
	err = rt.db.RemoveComment(dbcomment)
	if errors.Is(err, database.ErrCommentDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// Add a comment to a photo
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment

	currentTime := time.Now()

	// Obtain the text of the comment and the userid of the person who write the comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.PhotoId, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.Date = currentTime.Format("2006-01-02 15:04:05")

	err = CheckAuthentication(r.Header.Get("Authorization"), comment.UserId)
	if errors.Is(err, database.ErrNotAuthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbcomment := comment.CommentFromApiToDatabase()
	comment.Id, err = rt.db.SetComment(dbcomment)
	if err != nil && !errors.Is(err, database.ErrElementIsAlreadyExist) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)

}

// Get the list of comments of a photo
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comments []database.Comment

	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CheckAuthentication(r.Header.Get("Authorization"), userid)
	if errors.Is(err, database.ErrNotAuthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	comments, err = rt.db.GetComments(photoid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(comments)

}
