package database

import (
	"errors"
)

var ErrFollowDoesNotExist = errors.New(`Follow does not exist`)
var ErrUserDoesNotExist = errors.New(`User does not exist`)
var ErrPhotoDoesNotExist = errors.New(`Photo does not exist`)
var ErrLikeDoesNotExist = errors.New(`Like does not exist`)
var ErrCommentDoesNotExist = errors.New(`Comment does not exist`)
var ErrNotAuthorized = errors.New(`User is not authorized`)
var ErrBanDoesNotExist = errors.New(`Ban does not exist`)
var ErrElementIsAlreadyExist = errors.New("This element is already present into the database")

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

type Like struct {
	PhotoId uint64 `json:"photoId"`
	UserId  uint64 `json:"userId"`
	OwnerId uint64 `json:"ownerId"`
}

type Ban struct {
	UserId       uint64 `json:"userId"`
	BannedUserId uint64 `json:"bannedUserId"`
}

type Follow struct {
	UserId         uint64 `json:"userId"`
	FollowedUserId uint64 `json:"followedUserId"`
}

type Photo struct {
	Id             uint64 `json:"photoId"`
	Url            string `json:"url"`
	Date           string `json:"date"`
	LikeCounter    uint64 `json:"likeCounter"`
	CommentCounter uint64 `json:"commentCounter"`
	UserId         uint64 `json:"userId"`
}

type Comment struct {
	Text    string `json:"text"`
	Id      uint64 `json:"id"`
	UserId  uint64 `json:"userId"`
	PhotoId uint64 `json:"photoId"`
	Date    string `json:"date"`
}

type Stream struct {
	UserId uint64  `json:"userId"`
	Photos []Photo `json:"photos"`
}

type Profile struct {
	Username         string  `json:"username"`
	Photos           []Photo `json:"photos"`
	PhotoCounter     uint64  `json:"photoCounter"`
	FollowerCounter  uint64  `json:"followerCounter"`
	FollowingCounter uint64  `json:"followingCounter"`
}
