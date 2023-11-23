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

func (db *appdbimpl) CheckCommentId(commentid uint64) (bool, error){
	ris, err:= db.c.Exec(`SELECT * FROM Comment WHERE Id=?`, commentid)

	numberrows, err := ris.RowsAffected()
	if err != nil{
		return false, err
	} else if numberrows == 0 {
		return false, nil
	}
	return true, nil
	
}

func (db *appdbimpl) SetComment( c Comment)  error{
	_,err:= db.c.Exec(`INSERT INTO Comment (id, text, userid, photoid) VALUES (?, ?, ?, ?)`, c.Id, c.Text, c.UserId, c.PhotoId)
	return err
}