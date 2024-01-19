package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login
	rt.router.POST("/session", rt.wrap(rt.doLogin)) // DONE

	// Ban actions
	rt.router.PUT("/user/:uid/ban/:banneduid", rt.wrap(rt.banUser))      // DONE
	rt.router.DELETE("/user/:uid/ban/:banneduid", rt.wrap(rt.unbanUser)) // DONE
	rt.router.GET("/user/:uid/ban/:banneduid", rt.wrap(rt.isBan))        // DONE

	// Follow actions
	rt.router.PUT("/user/:uid/follow/:followeduid", rt.wrap(rt.followUser))      // DONE
	rt.router.DELETE("/user/:uid/follow/:followeduid", rt.wrap(rt.unfollowUser)) // DONE
	rt.router.GET("/user/:uid/follow/:followeduid", rt.wrap(rt.isFollow))        // DONE

	// User Information actions
	rt.router.PUT("/user/:uid/myusername", rt.wrap(rt.setMyUserName)) // DONE
	rt.router.GET("/user/:uid/stream", rt.wrap(rt.getMyStream))       // DONE
	rt.router.GET("/user/:uid/profile", rt.wrap(rt.getUserProfile))   // DONE
	rt.router.GET("/user", rt.wrap(rt.getUserBySubstring))            // DONE

	// Photo actions
	rt.router.DELETE("/user/:uid/photo/:photoid", rt.wrap(rt.deletePhoto)) // DONE
	rt.router.POST("/user/:uid/photo", rt.wrap(rt.uploadPhoto))            // DONE
	rt.router.GET("/user/:uid/photo", rt.wrap(rt.getPhotos))               // DONE

	// Like actions
	rt.router.PUT("/user/:uid/photo/:photoid/likes/:likeuid", rt.wrap(rt.likePhoto))      // DONE
	rt.router.DELETE("/user/:uid/photo/:photoid/likes/:likeuid", rt.wrap(rt.unlikePhoto)) // DONE
	rt.router.GET("/user/:uid/photo/:photoid/likes/:likeuid", rt.wrap(rt.isLike))         // DONE

	// Comment actions
	rt.router.DELETE("/user/:uid/photo/:photoid/comments/:commentid", rt.wrap(rt.uncommentPhoto)) // DONE
	rt.router.POST("/user/:uid/photo/:photoid/comments", rt.wrap(rt.commentPhoto))                // DONE
	rt.router.GET("/user/:uid/photo/:photoid/comments", rt.wrap(rt.getComments))                  // DONE

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
