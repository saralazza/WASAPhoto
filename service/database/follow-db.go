package database

func (db *appdbimpl) SetFollow(f Follow) error{
	_,err:= db.c.Exec(`INSERT INTO Follow (followeduserid, userid) VALUES (?, ?)`, f.FollowedUserId, f.UserId)
	if err != nil{
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveFollow(f Follow) error{
	ris, err := db.c.Exec(`DELETE FROM Follow WHERE followeduserid=? and userid=?`, f.FollowedUserId, f.UserId)
	if err != nil{
		return err
	}

	// Check if the follow exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return ErrorFollowDoesNotExist
	}
	return err
}
