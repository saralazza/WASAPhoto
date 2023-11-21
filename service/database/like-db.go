package database

func (db *appdbimpl) SetLike(f Like) error{
	_,err:= db.c.Exec(`INSERT INTO Like (photoid, userid) VALUES (?, ?)`, f.PhotoId, f.UserId)
	return err
}