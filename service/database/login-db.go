package database

import(
	"database/sql"
)

func (db *appdbimpl) SetUser(u User) error{
	_, err:= db.c.Exec(`INSERT INTO User (id, username) VALUES (?, ?)`, u.Id, u.Username)
	return err
}

func (db *appdbimpl) CheckUserId(userid uint64) (bool, error){
	err:= db.c.QueryRow(`SELECT Id FROM User WHERE Id=?`, userid).Scan()

	if err == sql.ErrNoRows{
		return false, nil
	} else if err != nil{
		return false, err
	}
	return true, nil
}