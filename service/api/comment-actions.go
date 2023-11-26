package api

import (
	"net/http"
	"strconv"
	"encoding/json"
	"time"
	"errors"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

/*
/user/{uid}/photo/{photoid}/comments/{commentid}:
    parameters:
      - {$ref: "#/components/parameters/uid"}
      - {$ref: "#/components/parameters/photoid"}
      - {$ref: "#/components/parameters/commentid"}
      
    delete:
      security:
      - bearerAuth: []
      description: Delete comment from a photo
      summary: Uncomment photo
      tags: ["comment"]
      operationId: uncommentPhoto
      responses:
        "204":
          description: Comment deleted successfully
        "400": {$ref: '#/components/responses/BadRequest'}
        "401": {$ref: '#/components/responses/Unauthorized'}
        "500": {$ref: '#/components/responses/InternalServerError'}*/

// Delete comment from a photo
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	var commentid uint64

	photoid, err := strconv.ParseUint(ps.ByName("photoid"),10,64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentid, err = strconv.ParseUint(ps.ByName("commentid"),10,64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.PhotoId = photoid
	comment.Id = commentid

	dbcomment := comment.CommentFromApiToDatabase()
	err = rt.db.RemoveComment(dbcomment)
	if errors.Is(err, database.ErrorCommentDoesNotExist){
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	// TODO : else if per l'errore sull'autorizzazione


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

	comment.PhotoId, err = strconv.ParseUint(ps.ByName("photoid"),10,64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.Date = currentTime.Format("2006-01-02 15:04:05")

	dbcomment := comment.CommentFromApiToDatabase()
	comment.Id,err = rt.db.SetComment(dbcomment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_=json.NewEncoder(w).Encode(comment)


}

// Get the list of comments of a photo
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}