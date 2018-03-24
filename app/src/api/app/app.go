package app

import (
	"api/app/items"
	"database/sql"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// Needed to sql lite 3
	_ "github.com/mattn/go-sqlite3"
)

var (
	r *gin.Engine
)

const (
	port string = ":8080"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	db := configDataBase()
	items.Configure(r, db)
	r.Run(port)
}

func configDataBase() *sql.DB {
	os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "user", "userpwd", "db", "db"))
	if err != nil {
		panic("Could not connect to the db")
	} 

	for {
		err := db.Ping()
		if err != nil {
			time.Sleep(1*time.Second)
			continue
		}
		// This is bad practice... You should create a schema.sql with all the definitions
		createTable(db)
		return db
	}

}

func createTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT
	);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
