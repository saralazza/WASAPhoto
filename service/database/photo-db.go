package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) SetPhoto(p Photo) (uint64,error){
	_,err:= db.c.Exec(`INSERT INTO Photo (userid, date, url) VALUES (?, ?, ?)`, p.UserId, p.Date, p.Url)
	if err != nil{
		return 0,err
	}
	
	var photoid uint64
	err = db.c.QueryRow(`SELECT id FROM Photo WHERE userid=? AND date=? AND url=?`, p.UserId, p.Date, p.Url).Scan(&photoid)
	if errors.Is(err, sql.ErrNoRows){
		return 0, ErrorPhotoDoesNotExist
	}else if err != nil{
		return 0,err
	}
	return photoid, nil
}

func (db *appdbimpl) RemovePhoto(p Photo)  error{
	ris, err := db.c.Exec(`DELETE FROM Photo WHERE id=?`, p.Id)
	if err != nil{
		return err
	}

	// Check if the photo exists really
	check, err := ris.RowsAffected()
	if err != nil{
		return err
	} else if check == 0 {
		return ErrorPhotoDoesNotExist
	}
	return err
}

func (db *appdbimpl) Stampa(){
	var photo Photo
	ris, _ := db.c.Query(`SELECT * FROM Photo`)

	for ris.Next(){
		err := ris.Scan(&photo.UserId,&photo.Id,&photo.Date,&photo.Url)

		if err !=nil{
			fmt.Printf("%s\n",err.Error())
		}

		fmt.Printf("userId %d photoid %d date %s url %s\n", photo.UserId, photo.Id,photo.Date,photo.Url)
	}



}