package database

import (
	"errors"
)

func (db *appdbimpl) SetPhoto(p Photo) error{
	_,err:= db.c.Exec(`INSERT INTO Photo (userid, date, url) VALUES (?, ?, ?)`, p.UserId, p.Date, p.Url)
	return err
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