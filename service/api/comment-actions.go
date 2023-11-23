package api

import (
	"net/http"
	"strconv"
	"encoding/json"
	"math/rand"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// Delete comment from a photo
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	var photoid uint64
	var commentid uint64
	var err error

	photoid, err = strconv.ParseUint(ps.ByName("photoid"),10,64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentid, err = strconv.ParseUint(ps.ByName("commentid"),10,64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment.PhotoId = photoid
	comment.Id = commentid

	var dbcomment database.Comment
	dbcomment = comment.CommentFromApiToDatabase()
	err = rt.db.RemoveComment(dbcomment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	
}

// Add a comment to a photo
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	var photoid uint64
	var commentid uint64
	var err error

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoid, err = strconv.ParseUint(ps.ByName("photoid"),10,64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentid = uint64(rand.Uint64())
	checkcommentid, err := rt.db.CheckCommentId(commentid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for checkcommentid {
		commentid = rand.Uint64()
		checkcommentid, err = rt.db.CheckCommentId(commentid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	comment.Id = commentid
	comment.PhotoId = photoid

	var dbcomment database.Comment
	dbcomment = comment.CommentFromApiToDatabase()
	err = rt.db.SetComment(dbcomment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)


}

// Get the list of comments of a photo
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}