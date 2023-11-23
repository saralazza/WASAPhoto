package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SetPhoto(p Photo) error{
	_,err:= db.c.Exec(`INSERT INTO Photo (id, userid, date, url) VALUES (?, ?, ?, ?)`, p.Id, p.UserId, p.Date, p.Url)
	return err
}

func (db *appdbimpl) CheckPhotoId(photoid uint64) (bool,error){
	err:= db.c.QueryRow(`SELECT Id FROM Photo WHERE Id=?`, photoid).Scan()

	if errors.Is(err, sql.ErrNoRows){
		return false, nil
	} else if err != nil{
		return false, err
	}
	return true, nil
}

func (db *appdbimpl) RemovePhoto(p Photo)  error{
	ris, err := db.c.Exec(`DELETE FROM Photo WHERE id=?`, p.Id)
	if err != nil{
		return err
	}

	// Check if the photo exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return errors.New(`Photo does not exist`)
	}
	return err
}