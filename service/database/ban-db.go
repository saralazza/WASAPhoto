package database

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) SetBan(b Ban) error {
	_, err := db.c.Exec(`INSERT INTO Ban (userid, banneduserid) VALUES (?, ?)`, b.UserId, b.BannedUserId)
	if err != nil {
		var sqlErr sqlite3.Error
		// Check if the tuple is already exist
		if errors.As(err, &sqlErr) && sqlErr.Code == sqlite3.ErrConstraint {
			// Chiave duplicata, gestisci di conseguenza
			return ErrElementIsAlreadyExist
		}

		return err
	}
	return nil
}

func (db *appdbimpl) RemoveBan(b Ban) error {
	ris, err := db.c.Exec(`DELETE FROM Ban WHERE banneduserid=? and userid=?`, b.BannedUserId, b.UserId)
	if err != nil {
		return err
	}

	// Check if the follow exists really
	check, err := ris.RowsAffected()
	if err != nil {
		return err
	} else if check == 0 {
		return ErrBanDoesNotExist
	}
	return nil

}
