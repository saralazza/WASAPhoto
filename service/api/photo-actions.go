package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"errors"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

)

// Delete a photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var photoid uint64

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photoid, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.Id = photoid
	photo.UserId = uid
	
	dbphoto := photo.PhotoFromApiToDatabase()
	err = rt.db.RemovePhoto(dbphoto)
	if errors.Is(err, database.ErrorPhotoDoesNotExist){
		http.Error(w, err.Error(), http.StatusFound)
	}
	// TODO : else if per l'errore sull'autorizzazione

	w.WriteHeader(http.StatusNoContent)
}

// Upload a photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var uid uint64

	currentTime := time.Now()

	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uid, err = strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.UserId = uid
	photo.Date = currentTime.Format("2006-01-02 15:04:05")
	photo.LikeCounter = 0
	photo.CommentCounter = 0

	dbphoto := photo.PhotoFromApiToDatabase()
	photo.Id, err = rt.db.SetPhoto(dbphoto)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_=json.NewEncoder(w).Encode(photo)

}

// Get the list of photos of an user
func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}