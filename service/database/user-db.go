package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SetUsername(u User) error {
	ris, err := db.c.Exec(`SELECT * FROM User WHERE username=?`, u.Username)
	if err != nil {
		return err
	}
	var check int64
	check, err = ris.RowsAffected()

	if err != nil {
		return err
	} else if check != 0 {
		return errors.New("Username must be unique")
	}

	ris, err = db.c.Exec(`UPDATE User SET username=? WHERE id=?`, u.Username, u.Id)
	if err != nil {
		return err
	}
	check, err = ris.RowsAffected()
	if err != nil {
		return err
	} else if check == 0 {
		return ErrorUserDoesNotExist
	}
	return nil
}

func (db *appdbimpl) GetStream(userid uint64) ([]Photo, error) {
	var ris []Photo

	rows, err := db.c.Query(`SELECT id, url, date FROM Photo WHERE userid IN 
		(SELECT followeduserid FROM Follow WHERE userid=? AND followeduserid NOT IN 
		(SELECT userid FROM Ban WHERE banneduserid=?))`, userid, userid)
	if err != nil {
		return nil, ErrorUserDoesNotExist
	}

	for rows.Next() {
		var photo Photo

		err = rows.Scan(&photo.Id, &photo.Url, &photo.Date)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(`SELECT userid FROM Photo WHERE id =?`, photo.Id).Scan(&photo.UserId)
		if err != nil {
			return nil, err
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		err = db.c.QueryRow(`SELECT COUNT(*) FROM Like WHERE photoid=?`, photo.Id).Scan(&photo.LikeCounter)
		if err != nil {
			return nil, err
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		err = db.c.QueryRow(`SELECT COUNT(*) FROM Comment WHERE photoid=?`, photo.Id).Scan(&photo.CommentCounter)
		if err != nil {
			return nil, err
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		ris = append(ris, photo)
	}

	_ = rows.Close()

	return ris, nil
}

func (db *appdbimpl) GetUsernameById(userid uint64) (string, error) {
	var username string
	err := db.c.QueryRow(`SELECT username FROM User WHERE id=?`, userid).Scan(&username)
	if err != nil {
		return "", err
	} else if errors.Is(err, sql.ErrNoRows) {
		return "", ErrorUserDoesNotExist
	}
	return username, nil
}

func (db *appdbimpl) GetProfile(userid uint64) (uint64, uint64, uint64, error) {
	var photoCounter uint64
	err := db.c.QueryRow(`SELECT COUNT(*) FROM Photo WHERE userid=?`, userid).Scan(&photoCounter)
	if err != nil {
		return 0, 0, 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, 0, ErrorUserDoesNotExist
	}

	var followerCounter uint64
	err = db.c.QueryRow(`SELECT COUNT(*) FROM Follow WHERE followeduserid=?`, userid).Scan(&followerCounter)
	if err != nil {
		return 0, 0, 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, 0, ErrorUserDoesNotExist
	}

	var followingCounter uint64
	err = db.c.QueryRow(`SELECT COUNT(*) FROM Follow WHERE userid=?`, userid).Scan(&followingCounter)
	if err != nil {
		return 0, 0, 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, 0, ErrorUserDoesNotExist
	}

	return photoCounter, followerCounter, followingCounter, nil
}
