package sqldb

import "database/sql"

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "username:password@/dbname")
	if err != nil {
		panic(err.Error())
	}

	return db
}
