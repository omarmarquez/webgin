package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

var msgChan chan string
var dbmap = initDb()

func index(c *gin.Context) {

	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func main() {
	// Main
	msgChan = make(chan string)
	fmt.Println("There you go.")

	app := gin.Default()
	app.GET("/", index)
	app.Run(":8000")

}

func initDb() {
	db, err = sql.Open("sqlite3", "db.sqlite3")
	checkErr(err, "initDb: Sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmpa.AddTableWithName(Article{}, "articles").SetKeys(true, "Id")
	checkErr(err, "initDb: Create tables failed")
	return dbmap
}

func checkErr(err error, loc, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
