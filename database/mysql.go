package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql(driver string, source string) *sql.DB {

	var db, err = sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
