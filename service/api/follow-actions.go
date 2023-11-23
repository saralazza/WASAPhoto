package api

import (
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"

	"strconv"
	"net/http"
	"encoding/json"
)

// Follow an user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	var followeuid uint64

	uid, err := strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followeuid, err = strconv.ParseUint(ps.ByName("followeduid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	follow.UserId = uid
	follow.FollowedUserId = followeuid

	dbFollow := follow.FollowFromApiToDatabase()
	err = rt.db.SetFollow(dbFollow)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_=json.NewEncoder(w).Encode(follow)

}

// Unfollow an user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	var uid uint64
	var followeuid uint64
	var err error

	uid, err= strconv.ParseUint(ps.ByName("uid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followeuid, err = strconv.ParseUint(ps.ByName("followeduid"), 10, 64)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	follow.UserId = uid
	follow.FollowedUserId = followeuid

	dbFollow := follow.FollowFromApiToDatabase()
	err = rt.db.RemoveFollow(dbFollow)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// Obtain the list of followed users
func (rt *_router) getFollowList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// TODO: come prendere lista?
}