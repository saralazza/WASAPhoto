package database

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) SetFollow(f Follow) error {
	_, err := db.c.Exec(`INSERT INTO Follow (followeduserid, userid) VALUES (?, ?)`, f.FollowedUserId, f.UserId)

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

func (db *appdbimpl) RemoveFollow(f Follow) error {
	ris, err := db.c.Exec(`DELETE FROM Follow WHERE followeduserid=? and userid=?`, f.FollowedUserId, f.UserId)
	if err != nil {
		return err
	}

	// Check if the follow exists really
	check, err := ris.RowsAffected()
	if err != nil {
		return err
	} else if check == 0 {
		return ErrFollowDoesNotExist
	}
	return err
}

func (db *appdbimpl) GetFollowings(userid uint64) ([]string, error) {
	var followings []string

	rows, err := db.c.Query(`SELECT followeduserid FROM Follow WHERE userid=?`, userid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uint64
		var following string

		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		following, err = db.GetUsernameById(id)
		if err != nil {
			return nil, err
		}

		followings = append(followings, following)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	_ = rows.Close()

	return followings, nil
}
