package api

import (
	"encoding/json"
	"errors"
	"io"
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
	photo.Username, err = rt.db.GetUsernameById(uid)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = CheckAuthentication(r.Header.Get("Authorization"), uid)
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

	currentTime := time.Now()

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.Url, err = io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.Username, err = rt.db.GetUsernameById(uid)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.Date = currentTime.Format("2006-01-02 15:04:05")
	photo.LikeCounter = 0
	photo.CommentCounter = 0
	photo.UserId = uid

	err = CheckAuthentication(r.Header.Get("Authorization"), uid)
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
	var username string

	userid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, err = rt.db.GetUsernameById(userid)
	if errors.Is(err, database.ErrUserDoesNotExist) {
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

	photos, err = rt.db.GetPhotos(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photos)
}
