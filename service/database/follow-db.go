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

/*func (db *appdbimpl) Stampa(){
	var follow Follow
	ris, _ := db.c.Query(`SELECT * FROM Follow`)

	for ris.Next(){
		_ = ris.Scan(&follow.UserId,&follow.FollowedUserId)

		fmt.Printf("userId %d followeduid %d\n", follow.UserId, follow.FollowedUserId)
	}



}*/
