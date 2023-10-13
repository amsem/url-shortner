package helper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
// "am53m*********"

const (
    host = "127.0.0.1"
    port = 5432
    user = "amsem"
    password = "am53m*********" 
    dbname = "mydb"
)

func InitDB() (*sql.DB, error) {
    var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    
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
