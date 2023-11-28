package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
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
		return ErrorCommentDoesNotExist
	}
	return err
}

func (db *appdbimpl) SetComment( c Comment)  (uint64,error){
	_,err:= db.c.Exec(`INSERT INTO Comment (text, userid, photoid, date) VALUES (?, ?, ?, ?)`, c.Text, c.UserId, c.PhotoId,c.Date)
	if err != nil{
		var sqlErr sqlite3.Error
		// Check if the tuple is already exist
		if errors.As(err, &sqlErr) && sqlErr.Code == sqlite3.ErrConstraint {
			// Chiave duplicata, gestisci di conseguenza
			return 0, ErrorElementIsAlreadyExist
		}
		
		return 0,err
	}
	
	var commentid uint64
	err = db.c.QueryRow(`SELECT id FROM Comment WHERE text=? AND userid=? AND photoid=? AND date=?`, c.Text, c.UserId, c.PhotoId,c.Date).Scan(&commentid)
	if errors.Is(err, sql.ErrNoRows){
		return 0, ErrorCommentDoesNotExist
	}else if err != nil{
		return 0,err
	}
	return commentid, nil
}

func (db *appdbimpl) ObtainCommentUserId( commentid uint64)  (uint64,error){
	var uid uint64
	err := db.c.QueryRow(`SELECT userid FROM Comment WHERE id=?`, commentid).Scan(&uid)
	if errors.Is(err, sql.ErrNoRows){
		return 0, ErrorCommentDoesNotExist
	}else if err != nil{
		return 0,err
	}
	return uid, nil
}

func (db *appdbimpl) RemoveComments( b Ban) error{
	_, err := db.c.Exec(`DELETE FROM Comment WHERE userid=? AND photoid IN (SELECT id FROM Photo WHERE userid=?)`,b.BannedUserId,b.UserId)
	if err != nil {
		return err
	}
	return nil
}
