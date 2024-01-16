package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"

	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Set username of the user
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var username string

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

	err = CheckAuthentication(r.Header.Get("Authorization"), user.Id)
	if errors.Is(err, database.ErrNotAuthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	username, err = rt.db.GetUsernameById(user.Id)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if username == user.Username {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbuser := user.UserFromApiToDatabase()
	err = rt.db.SetUsername(dbuser)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// Get the user stream composed by photos from following users
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	photos, err = rt.db.GetStream(userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photos)

}

// Get user profile composed by the user’s photos, how many photos have been uploaded, and the user’s followers and following.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var profile database.Profile

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

	profile.Username, err = rt.db.GetUsernameById(userid)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.PhotoCounter, profile.FollowerCounter, profile.FollowingCounter, err = rt.db.GetProfile(userid)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(profile)

}

// Search users which username contains a substring specified into the path url
func (rt *_router) getUserBySubstring(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var users []database.User

	substring := r.URL.Query().Get("substring")

	users, err := rt.db.SearchUsers(substring)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)
}
