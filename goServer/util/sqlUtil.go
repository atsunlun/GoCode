package util

import (
    "crypto/md5"
    "database/sql"
    "encoding/hex"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
    "time"
)

const (
	//user:password@tcp(localhost:5555)/dbname?charset=utf8
	DB_Driver = "root:777465@tcp(49.234.183.79:3306)/server_db?charset=utf8"
)

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	CheckErr(err)
	return isOpen, db
}
