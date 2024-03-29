package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SetUsername(u User) error {

	ris, err := db.c.Exec(`UPDATE User SET username=? WHERE id=?`, u.Username, u.Id)
	if err != nil {
		return err
	}
	check, err := ris.RowsAffected()
	if err != nil {
		return err
	} else if check == 0 {
		return ErrUserDoesNotExist
	}

	ris, err = db.c.Exec(`UPDATE Photo SET username=? WHERE uid=?`, u.Username, u.Id)
	if err != nil {
		return err
	}
	check, err = ris.RowsAffected()
	if err != nil {
		return err
	} else if check == 0 {
		return ErrUserDoesNotExist
	}
	return nil
}

func (db *appdbimpl) GetStream(userid uint64) ([]Photo, error) {
	var ris []Photo

	rows, err := db.c.Query(`SELECT * FROM Photo WHERE uid IN 
		(SELECT followeduserid FROM Follow WHERE userid=? AND followeduserid NOT IN 
		(SELECT userid FROM Ban WHERE banneduserid=?))`, userid, userid)
	if err != nil {
		return nil, ErrUserDoesNotExist
	}

	for rows.Next() {
		var photo Photo

		err = rows.Scan(&photo.Id, &photo.Username, &photo.Date, &photo.Url, &photo.UserId)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow(`SELECT username FROM Photo WHERE id =?`, photo.Id).Scan(&photo.Username)
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

	if err := rows.Err(); err != nil {
		return nil, err
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
		return "", ErrUserDoesNotExist
	}
	return username, nil
}

func (db *appdbimpl) GetProfile(userid uint64) (uint64, uint64, uint64, error) {
	var photoCounter uint64
	err := db.c.QueryRow(`SELECT COUNT(*) FROM Photo WHERE uid=?`, userid).Scan(&photoCounter)
	if err != nil {
		return 0, 0, 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, 0, ErrUserDoesNotExist
	}

	var followerCounter uint64
	err = db.c.QueryRow(`SELECT COUNT(*) FROM Follow WHERE followeduserid=?`, userid).Scan(&followerCounter)
	if err != nil {
		return 0, 0, 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, 0, ErrUserDoesNotExist
	}

	var followingCounter uint64
	err = db.c.QueryRow(`SELECT COUNT(*) FROM Follow WHERE userid=?`, userid).Scan(&followingCounter)
	if err != nil {
		return 0, 0, 0, err
	} else if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, 0, ErrUserDoesNotExist
	}

	return photoCounter, followerCounter, followingCounter, nil
}

func (db *appdbimpl) SearchUsers(substring string, uid uint64) ([]User, error) {
	var users []User

	rows, err := db.c.Query(`SELECT * FROM User WHERE id IN( SELECT id FROM User WHERE username LIKE '%'||?||'%' EXCEPT SELECT userid FROM Ban WHERE banneduserid=? EXCEPT SELECT ?)`, substring, uid, uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	_ = rows.Close()

	return users, nil

}
