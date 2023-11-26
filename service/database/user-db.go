package database

import(
	"errors"
)

func (db *appdbimpl) SetUsername(u User) error{
	ris, err := db.c.Exec(`SELECT * FROM User WHERE username=?`, u.Username)
	if err != nil{
		return err
	}
	var check int64
	check, err = ris.RowsAffected()

	if err != nil{
		return err
	}else if check!=0{
		return errors.New("Username must be unic")
	}

	ris, err = db.c.Exec(`UPDATE User SET username=? WHERE id=?`, u.Username,u.Id)
	if err != nil{
		return err
	}
	check, err = ris.RowsAffected()
	if err != nil{
		return err
	}else if check == 0{
		return ErrorUserDoesNotExist
	}
	return nil
}