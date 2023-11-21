package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"math/rand"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"

)

// Delete a photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var uid uint64
	var photoid uint64
	var err error

	uid, err = strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoid, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.Id = photoid
	photo.UserId = uid
	
	var dbphoto database.Photo 
	dbphoto = photo.PhotoFromApiToDatabase()
	err = rt.db.RemovePhoto(dbphoto)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Upload a photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var uid uint64
	var url string
	var photoid uint64
	var err error

	currentTime := time.Now()

	uid, err = strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url = ps.ByName("url")

	photoid = uint64(rand.Uint64())
	checkphotoid, err := rt.db.CheckPhotoId(photoid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for checkphotoid {
		photoid = rand.Uint64()
		checkphotoid, err = rt.db.CheckPhotoId(photoid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	photo.Id = photoid
	photo.Url = url
	photo.Date = currentTime.Format("2006-01-02 15:04:05")
	photo.LikeCounter = 0
	photo.CommentCounter = 0
	photo.UserId = uid

	var dbphoto database.Photo
	dbphoto = photo.PhotoFromApiToDatabase()
	err = rt.db.SetPhoto(dbphoto)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(photo)

}

// Get the list of photos of an user
func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}