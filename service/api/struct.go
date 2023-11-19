package api

import(
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type User struct{
	Id uint64 `json:"id"`
	Username string `json:"username"`
}

type Like struct{
	PhotoId uint64 `json:"photoId"`
	UserId uint64 `json:"userId"`
	OwnerId uint64 `json:"ownerId"`
}

type Ban struct{
	UserId uint64 `json:"userId"`
	BannedUserId uint64 `json:"bannedUserId"`
}

type Follow struct{
	UserId uint64 `json:"userId"`
	FollowedUserId uint64 `json:"followedUserId"`
}

type Photo struct{
	Id uint64 `json:"photoId"`
	Url string `json:"url"`
	Date string `json:"date"`
	LikeCounter uint64 `json:"likeCounter"`
	CommentCounter uint64 `json:"commentCounter"`
	UserId uint64 `json:"userId"`
}

type Comment struct{
	Text string `json:"text"`
	Id uint64 `json:"id"`
	UserId uint64 `json:"userId"`
}