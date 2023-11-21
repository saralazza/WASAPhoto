//SetPhoto(Photo) error
//GetLastPhotoId() (int,error)

package database

func (db *appdbimpl) SetPhoto(p Photo) error{
	_,err:= db.c.Exec(`INSERT INTO Photo (id, userid, date, url) VALUES (?, ?, ?, ?)`, p.Id, p.UserId, p.Date, p.Url)
	return err
}

func (db *appdbimpl) CheckPhotoId(photoid uint64) (bool, error){
	ris, err:= db.c.Exec(`SELECT * FROM Photo WHERE Id=?`, photoid)

	numberrows, err := ris.RowsAffected()
	if err != nil{
		return false, err
	} else if numberrows == 0 {
		return false, nil
	}
	return true, nil
}