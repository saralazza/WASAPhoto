package database

func (db *appdbimpl) SetBan(b Ban)error{
	_,err:= db.c.Exec(`INSERT INTO Ban (userid, banneduserid) VALUES (?, ?)`, b.UserId, b.BannedUserId)
	if err != nil{
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveBan(b Ban)error{
	ris, err := db.c.Exec(`DELETE FROM Ban WHERE bannededuserid=? and userid=?`, b.BannedUserId, b.UserId)
	if err != nil{
		return err
	}

	// Check if the follow exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return ErrorBanDoesNotExist
	}
	return err
	
}