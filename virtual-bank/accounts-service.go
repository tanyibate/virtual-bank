package virtualBank

import (
	"database/sql"
	"github.com/lib/pq"
)

var connStr string = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
var db, dbErr = sql.Open("postgres", connStr)

func addUserToDB(account *Account, db *sql.DB) {

}
