package db

import (
	"database/sql"

	"filec2/config"
)

var (
	//SQL is a SQL DB
	SQL *sql.DB
)

//Configure will connect to the database
func Configure() error {
	db, err := sql.Open("mysql", config.SQLUser+":"+config.SQLPassword+"@tcp("+config.SQLHost+")/"+config.SQLDatabase)
	if err != nil {
		return err
	}

	SQL = db
	return nil
}
