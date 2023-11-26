package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login
	rt.router.POST("/session",rt.wrap(rt.doLogin)) // DONE

	// Ban actions
	rt.router.PUT("/user/:uid/ban/:banneduid",rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:uid/ban/:banneduid",rt.wrap(rt.unbanUser))
	rt.router.GET("/user/:uid/ban",rt.wrap(rt.getBanList))

	// Follow actions
	rt.router.PUT("/user/:uid/follow/:followeduid",rt.wrap(rt.followUser)) // DONE
	rt.router.DELETE("/user/:uid/follow/:followeduid",rt.wrap(rt.unfollowUser)) // DONE
	rt.router.GET("/user/:uid/follow",rt.wrap(rt.getFollowList))

	// User Information actions
	rt.router.PUT("/user/:uid/myusername",rt.wrap(rt.setMyUserName))
	rt.router.GET("/user/:uid/stream",rt.wrap(rt.getMyStream))
	rt.router.GET("/user/:uid/profile",rt.wrap(rt.getUserProfile))

	// Photo actions
	rt.router.DELETE("/user/:uid/photo/:photoid",rt.wrap(rt.deletePhoto))
	rt.router.POST("/user/:uid/photo",rt.wrap(rt.uploadPhoto))
	rt.router.GET("/user/:uid/photo",rt.wrap(rt.getPhotos))

	// Like actions
	rt.router.PUT("/user/:uid/photo/:photoid/likes/:likeuid",rt.wrap(rt.likePhoto)) 
	rt.router.DELETE("/user/:uid/photo/:photoid/likes/:likeuid",rt.wrap(rt.unlikePhoto)) 
	rt.router.GET("/user/:uid/photo/:photoid/likes",rt.wrap(rt.getLikes))

	// Comment actions
	rt.router.DELETE("/user/:uid/photo/:photoid/comments/:commentid",rt.wrap(rt.uncommentPhoto)) 
	rt.router.POST("/user/:uid/photo/:photoid/comments",rt.wrap(rt.commentPhoto))
	rt.router.GET("/user/:uid/photo/:photoid/comments",rt.wrap(rt.getComments))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
