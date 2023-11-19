package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// If the user does not exist, it will be created, and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}
