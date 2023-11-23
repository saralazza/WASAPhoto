package api

import (
	"net/http"
	"strconv"

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

}

// Get the list of comments of a photo
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}