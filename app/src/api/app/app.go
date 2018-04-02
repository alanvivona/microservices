package app

import (
	"api/app/gdrive"
	"api/app/items"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

import _ "github.com/go-sql-driver/mysql" // Import needed fo mysql drivers

var (
	r *gin.Engine
)

const (
	appPort string = ":8080"
	dbPort  string = "3306"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	db := configDataBase()
	items.Configure(r, db)
	gdrive.Configure(r)
	r.Run(appPort)
}

func configDataBase() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "user", "userpwd", "db", dbPort, "db"))
	if err != nil {
		panic("Could not connect to the db")
	} else {
		return db
	}
}
