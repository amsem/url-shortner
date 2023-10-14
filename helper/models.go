package helper

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

)

var connectString = os.Getenv("CONNSTR")

func InitDB() (*sql.DB, error) {

    db, err := sql.Open("postgres", connectString)
    if err != nil {
        return nil, err
    }
    stmt, e := db.Prepare("CREATE TABLE IF NOT EXISTS web_url (ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")

    if e != nil {
        return nil, e
    }
    _, err = stmt.Exec()

    if err != nil {
        return nil, err
    }
    return db, nil
}
