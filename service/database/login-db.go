package database

import(
	"database/sql"
	"errors"
)

func (db *appdbimpl) SetUser(u User) error{
	_, err:= db.c.Exec(`INSERT INTO User (id, username) VALUES (?, ?)`, u.Id, u.Username)
	return err
}

func (db *appdbimpl) CheckUsername(username string) (uint64, error){
	var id uint64
	err:= db.c.QueryRow(`SELECT Id FROM User WHERE Username=?`, username).Scan(&id)

	if err != nil{
		return 0, err
	}
	return id, nil
}

func (db *appdbimpl) CheckUserId(userid uint64) (bool, error){
	err:= db.c.QueryRow(`SELECT Id FROM User WHERE Id=?`, userid).Scan()

	if errors.Is(err, sql.ErrNoRows){
		return false, nil
	} else if err != nil{
		return false, err
	}
	return true, nil
}