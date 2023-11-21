package database

func (db *appdbimpl) SetFollow(f Follow) error{
	_,err:= db.c.Exec(`INSERT INTO Follower (followeduserid, userid) VALUES (?, ?)`, f.FollowedUserId, f.UserId)
	return err
}