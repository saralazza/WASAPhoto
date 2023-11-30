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

// Delete a photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var photoid uint64

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photoid, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.Id = photoid
	photo.UserId = uid

	err = CheckAuthentication(r.Header.Get("Authorization"), photo.UserId)
	if errors.Is(err, database.ErrNotAuthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbphoto := photo.PhotoFromApiToDatabase()
	err = rt.db.RemovePhoto(dbphoto)
	if errors.Is(err, database.ErrPhotoDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.UserId = uid
	photo.Date = currentTime.Format("2006-01-02 15:04:05")
	photo.LikeCounter = 0
	photo.CommentCounter = 0

	err = CheckAuthentication(r.Header.Get("Authorization"), photo.UserId)
	if errors.Is(err, database.ErrNotAuthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbphoto := photo.PhotoFromApiToDatabase()
	photo.Id, err = rt.db.SetPhoto(dbphoto)
	if err != nil && !errors.Is(err, database.ErrElementIsAlreadyExist) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)

}

// Get the list of photos of an user
func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photos []database.Photo

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

	photos, err = rt.db.GetPhotos(userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photos)
}
