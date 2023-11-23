package database

func (db *appdbimpl) SetUser(u User) error{
	_, err:= db.c.Exec(`INSERT INTO User (id, username) VALUES (?, ?)`, u.Id, u.Username)
	return err
}

func (db *appdbimpl) CheckUserId(userid uint64) (bool, error){
	ris, err:= db.c.Exec(`SELECT * FROM User WHERE Id=?`, userid)

	numberrows, err := ris.RowsAffected()
	if err != nil{
		return false, err
	} else if numberrows == 0 {
		return false, nil
	}
	return true, nil
}