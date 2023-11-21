package database

import(
	"errors"
)

func (db *appdbimpl) SetLike(l Like) error{
	_,err:= db.c.Exec(`INSERT INTO Like (photoid, userid) VALUES (?, ?)`, l.PhotoId, l.UserId)
	return err
}

func (db *appdbimpl) RemoveLike(l Like) error{
	ris, err := db.c.Exec(`DELETE FROM Like WHERE photoid=? and userid=?`, l.PhotoId, l.UserId)
	if err != nil{
		return err
	}

	// Check if the follow exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return errors.New(`Like does not exist`)
	}
	return err
}