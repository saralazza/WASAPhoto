package database

import(
	"errors"
)

func (db *appdbimpl) RemoveComment(c Comment) error{
	ris, err := db.c.Exec(`DELETE FROM Comment WHERE id=?`, c.Id)
	if err != nil{
		return err
	}

	// Check if the follow exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return errors.New(`Comment does not exist`)
	}
	return err
}