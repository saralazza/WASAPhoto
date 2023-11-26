package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Follow an user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	var followeuid uint64

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	followeuid, err = strconv.ParseUint(ps.ByName("followeduid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	follow.UserId = uid
	follow.FollowedUserId = followeuid

	err = CheckAuthentication(r.Header.Get("Authorization"),uid)
	if errors.Is(err,database.ErrorNotAuthorized){
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbFollow := follow.FollowFromApiToDatabase()
	err = rt.db.SetFollow(dbFollow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(follow)

}

// Unfollow an user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	var uid uint64
	var followeuid uint64
	var err error

	uid, err = strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	followeuid, err = strconv.ParseUint(ps.ByName("followeduid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	follow.UserId = uid
	follow.FollowedUserId = followeuid

	err = CheckAuthentication(r.Header.Get("Authorization"),uid)
	if errors.Is(err,database.ErrorNotAuthorized){
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbFollow := follow.FollowFromApiToDatabase()
	err = rt.db.RemoveFollow(dbFollow)
	if errors.Is(err, database.ErrorFollowDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	// TODO : else if per l'errore sull'autorizzazione

	w.WriteHeader(http.StatusNoContent)

}

// Obtain the list of followed users
func (rt *_router) getFollowList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// TODO: come prendere lista?
}
