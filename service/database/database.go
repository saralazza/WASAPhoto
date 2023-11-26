/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// TODO: cambiare nome
// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error

	SetFollow(Follow) error
	RemoveFollow(Follow) error

	SetLike(Like) error
	RemoveLike(Like) error
	RemoveLikes(Ban)error

	SetPhoto(Photo) (uint64, error)
	RemovePhoto(Photo) error
	GetPhotos(uint64)([]Photo,error)

	SetUser(User) (uint64,error)
	CheckUsername(string)(uint64,error)
	SetUsername(User)error
	GetStream(uint64)([]Photo, error)
	GetUsernameById(uint64)(string, error)
	GetProfile(uint64)(uint64,uint64,uint64,error)

	RemoveComment(Comment) error
	SetComment(Comment) (uint64,error)
	ObtainCommentUserId(uint64) (uint64, error)
	RemoveComments(Ban)error

	SetBan(Ban)error
	RemoveBan(Ban)error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	_,err := db.Exec("PRAGMA foreign_key = ON")
	if err != nil {
		return nil, err
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	sqlStmt := `CREATE TABLE IF NOT EXISTS User (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		username TEXT  NOT NULL UNIQUE
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Photo (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, userid TEXT NOT NULL, 
		date TEXT NOT NULL, url TEXT NOT NULL,
		FOREIGN KEY (userid) REFERENCES User(id)
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Like (photoid INTEGER NOT NULL, userid TEXT  NOT NULL, 
		PRIMARY KEY (photoid, userid),
		FOREIGN KEY (photoid) REFERENCES Photo(id),
		FOREIGN KEY (userid) REFERENCES User(id)
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Ban (userid INTEGER NOT NULL, banneduserid TEXT  NOT NULL, 
		PRIMARY KEY (banneduserid, userid),
		FOREIGN KEY (userid) REFERENCES User(id),
		FOREIGN KEY (banneduserid) REFERENCES User(id)
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Follow (userid INTEGER NOT NULL, followeduserid TEXT  NOT NULL, 
		PRIMARY KEY (followeduserid, userid),
		FOREIGN KEY (userid) REFERENCES User(id),
		FOREIGN KEY (followeduserid) REFERENCES User(id)
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Comment (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, text TEXT  NOT NULL, 
		userid INTEGER NOT NULL, photoid INTEGER NOT NULL, date TEXT NOT NULL,
		FOREIGN KEY (userid) REFERENCES User(id)
		FOREIGN KEY (photoid) REFERENCES Photo(id)
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
