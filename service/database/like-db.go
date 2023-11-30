package database

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) SetLike(l Like) error {
	_, err := db.c.Exec(`INSERT INTO Like (photoid, userid) VALUES (?, ?)`, l.PhotoId, l.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveLike(l Like) error {
	ris, err := db.c.Exec(`DELETE FROM Like WHERE photoid=? and userid=?`, l.PhotoId, l.UserId)
	if err != nil {
		var sqlErr sqlite3.Error
		// Check if the tuple is already exist
		if errors.As(err, &sqlErr) && sqlErr.Code == sqlite3.ErrConstraint {
			// Chiave duplicata, gestisci di conseguenza
			return ErrElementIsAlreadyExist
		}

		return err
	}

	// Check if the like exists really
	check, err := ris.RowsAffected()
	if err != nil {
		return err
	} else if check == 0 {
		return ErrLikeDoesNotExist
	}
	return err
}

func (db *appdbimpl) RemoveLikes(b Ban) error {
	_, err := db.c.Exec(`DELETE FROM Like WHERE userid=? AND photoid IN (SELECT id FROM Photo WHERE userid=?)`, b.BannedUserId, b.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetLikes(photoid uint64) ([]string, error) {
	var likes []string

	rows, err := db.c.Query(`SELECT userid FROM Like WHERE photoid=?`, photoid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uint64
		var like string

		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		like, err = db.GetUsernameById(id)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	if err := rows.Err(); err != nil {
        return nil, err
    }

	_ = rows.Close()

	return likes, nil
}
