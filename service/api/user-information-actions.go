package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// Set username of the user
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}

// Get the user stream composed by photos from following users
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}

// Get user profile composed by the user’s photos, how many photos have been uploaded, and the user’s followers and following.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}