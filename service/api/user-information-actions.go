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

	dbuser := user.UserFromApiToDatabase()
	err = rt.db.SetUsername(dbuser)
	if errors.Is(err, database.ErrorUserDoesNotExist){
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}else if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	} // TODO aggiungere errore autorizzazione

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}


// Get the user stream composed by photos from following users
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}

// Get user profile composed by the user’s photos, how many photos have been uploaded, and the user’s followers and following.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}