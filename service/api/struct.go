package api

import(
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type User struct{
	Id uint64 `json:"id"`
	Username string `json:"username"`
}

func UserFromDatabaseToApi(user database.User) User{
	u := User{
		Id: user.Id,
		Username: user.Username,
	}
	
	return u
}

func (u *User) UserFromApiToDatabase() database.User{
	user := database.User{
		Id: u.Id,
		Username: u.Username,
	}

	return user
}

type Like struct{
	PhotoId uint64 `json:"photoId"`
	UserId uint64 `json:"userId"`
	OwnerId uint64 `json:"ownerId"`
}

func LikeFromDatabaseToApi(like database.Like)Like{
	l := Like{
		PhotoId: like.PhotoId,
		UserId: like.UserId,
		OwnerId: like.OwnerId,
	}
	
	return l
}

func (l *Like) LikeFromApiToDatabase() database.Like{
	like := database.Like{
		PhotoId: l.PhotoId,
		UserId: l.UserId,
		OwnerId: l.OwnerId,
	}

	return like
}

type Ban struct{
	UserId uint64 `json:"userId"`
	BannedUserId uint64 `json:"bannedUserId"`
}

func BanFromDatabaseToApi(ban database.Ban) Ban{
	b := Ban{
		UserId: ban.UserId,
		BannedUserId: ban.BannedUserId,
	}
	
	return b
}

func (b *Ban) BanFromApiToDatabase() database.Ban{
	ban := database.Ban{
		UserId: b.UserId,
		BannedUserId: b.BannedUserId,
	}

	return ban
}

type Follow struct{
	UserId uint64 `json:"userId"`
	FollowedUserId uint64 `json:"followedUserId"`
}

func FollowFromDatabaseToApi(follow database.Follow) Follow{
	f := Follow{
		UserId: follow.UserId,
		FollowedUserId: follow.FollowedUserId,
	}
	
	return f
}

func (f *Follow) FollowFromApiToDatabase() database.Follow{
	follow := database.Follow{
		UserId: f.UserId,
		FollowedUserId: f.FollowedUserId,
	}

	return follow
}

type Photo struct{
	Id uint64 `json:"photoId"`
	Url string `json:"url"`
	Date string `json:"date"`
	LikeCounter uint64 `json:"likeCounter"`
	CommentCounter uint64 `json:"commentCounter"`
	UserId uint64 `json:"userId"`
}

func PhotoFromDatabaseToApi(photo database.Photo) Photo{
	p := Photo{
		Id: photo.Id,
		Url: photo.Url,
		Date: photo.Date,
		LikeCounter: photo.LikeCounter,
		CommentCounter: photo.CommentCounter,
		UserId: photo.UserId,
	}
	
	return p
}

func (p *Photo) PhotoFromApiToDatabase() database.Photo{
	photo := database.Photo{
		Id: p.Id,
		Url: p.Url,
		Date: p.Date,
		LikeCounter: p.LikeCounter,
		CommentCounter: p.CommentCounter,
		UserId: p.UserId,
	}

	return photo
}

type Comment struct{
	Text string `json:"text"`
	Id uint64 `json:"id"`
	UserId uint64 `json:"userId"`
	PhotoId uint64 `json:"photoId"`
	Date string `json:"date"` 
}

func CommentFromDatabaseToApi(comment database.Comment) Comment{
	c := Comment{
		Text: comment.Text,
		Id: comment.Id,
		UserId: comment.UserId,
		PhotoId: comment.PhotoId,
		Date: comment.Date,
	}
	
	return c
}

func (c *Comment) CommentFromApiToDatabase() database.Comment{
	comment := database.Comment{
		Text: c.Text,
		Id: c.Id,
		UserId: c.UserId,
		PhotoId: c.PhotoId,
		Date: c.Date,
	}

	return comment
}