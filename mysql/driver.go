package mysql

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Init() {
	var err error
	Db, err = sql.Open("mysql", "banboom:11111111@/banboom")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}

func Close() {
	Db.Close()
}

func SelectAll() {
	results, err := Db.Query("SELECT * FROM test_table")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}
}