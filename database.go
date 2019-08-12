package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))

func getPaste(id uint64) (paste, title, author, notes []byte, err error) {
	row := db.QueryRow("SELECT paste, title, author, notes FROM pastes WHERE id = ?", id)
	err = row.Scan(&paste, &title, &author, &notes)
	return
}

func postPaste(paste, title, author, notes *string) (id uint64, err error) {
	if len(*title) == 0 {
		title = nil
	}

	if len(*author) == 0 {
		author = nil
	}

	if len(*notes) == 0 {
		notes = nil
	}

	res, err := db.Exec("INSERT INTO pastes(paste, title, author, notes) VALUES(?, ?, ?, ?)", paste, title, author, notes)
	if err != nil {
		return
	}

	idSigned, err := res.LastInsertId()
	if err != nil {
		return
	}

	id = uint64(idSigned)
	return
}
