package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"database/sql"
)

// Set username of the user
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Id, err = strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CheckAuthentication(r.Header.Get("Authorization"),user.Id)
	if errors.Is(err,database.ErrorNotAuthorized){
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbuser := user.UserFromApiToDatabase()
	err = rt.db.SetUsername(dbuser)
	if errors.Is(err, database.ErrorUserDoesNotExist){
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}else if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}


// Get the user stream composed by photos from following users
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var stream database.Stream

	userid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stream.UserId = userid

	err = CheckAuthentication(r.Header.Get("Authorization"),userid)
	if errors.Is(err,database.ErrorNotAuthorized){
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	stream.Photos, err = rt.db.GetStream(userid)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(stream)
	
}

// Get user profile composed by the user’s photos, how many photos have been uploaded, and the user’s followers and following.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var profile database.Profile

	userid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CheckAuthentication(r.Header.Get("Authorization"),userid)
	if errors.Is(err,database.ErrorNotAuthorized){
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	profile.Username, err = rt.db.GetUsernameById(userid)
	if errors.Is(err,sql.ErrNoRows){
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}else if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.Photos, err = rt.db.GetPhotos(userid)
	if errors.Is(err,database.ErrorUserDoesNotExist){
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}else if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.PhotoCounter, profile.FollowerCounter, profile.FollowingCounter, err= rt.db.GetProfile(userid)
	if errors.Is(err,database.ErrorUserDoesNotExist){
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}else if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(profile)
	
}