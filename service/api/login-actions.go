package api

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"errors"
	"database/sql"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// If the user does not exist, it will be created, and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User 
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id, err = rt.db.CheckUsername(user.Username)
	if errors.Is(err, sql.ErrNoRows){
		userid := uint64(rand.Int())
		var checkuserid bool
		checkuserid, err = rt.db.CheckUserId(userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for checkuserid {
			userid = uint64(rand.Int())
			checkuserid, err = rt.db.CheckUserId(userid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		user.Id = userid
		
		dbuser := user.UserFromApiToDatabase()
		err = rt.db.SetUser(dbuser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_=json.NewEncoder(w).Encode(user)
}