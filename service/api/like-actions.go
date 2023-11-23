package api

import (
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"

	"strconv"
	"net/http"
	"encoding/json"
)

// Add like to a photo
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	var likeuid uint64
	var photoid uint64

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	likeuid, err = strconv.ParseUint(ps.ByName("likeuid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoid, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	like.UserId = likeuid
	like.OwnerId = uid
	like.PhotoId = photoid
	
	dbLike := like.LikeFromApiToDatabase()
	err = rt.db.SetLike(dbLike)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_=json.NewEncoder(w).Encode(like)

}

// Delete like from a photo 
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	var likeuid uint64
	var photoid uint64

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	likeuid, err= strconv.ParseUint(ps.ByName("likeuid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoid, err= strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	like.UserId = likeuid
	like.OwnerId = uid
	like.PhotoId = photoid

	dbLike := like.LikeFromApiToDatabase()
	err = rt.db.RemoveLike(dbLike)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// Get the list of likes of a photo
func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}