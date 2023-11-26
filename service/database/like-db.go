package database

func (db *appdbimpl) SetLike(l Like) error{
	_,err:= db.c.Exec(`INSERT INTO Like (photoid, userid) VALUES (?, ?)`, l.PhotoId, l.UserId)
	if err != nil{
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveLike(l Like) error{
	ris, err := db.c.Exec(`DELETE FROM Like WHERE photoid=? and userid=?`, l.PhotoId, l.UserId)
	if err != nil{
		return err
	}

	// Check if the like exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return ErrorLikeDoesNotExist
	}
	return err
}