package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SetUser(u User) (uint64, error) {
	_, err := db.c.Exec(`INSERT INTO User (username) VALUES (?)`, u.Username)
	if err == nil {
		var userid uint64
		err = db.c.QueryRow(`SELECT id FROM User WHERE Username=?`, u.Username).Scan(&userid)
		if errors.Is(err, sql.ErrNoRows) {
			return userid, ErrUserDoesNotExist
		}
		return userid, nil
	}
	return 0, err
}

func (db *appdbimpl) CheckUsername(username string) (uint64, error) {
	var id uint64
	err := db.c.QueryRow(`SELECT id FROM User WHERE Username=?`, username).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
