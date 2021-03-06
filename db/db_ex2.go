package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func DBOne() {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://jqnqjfrd:%s@%s/jqnqjfrd",
		os.Getenv("ELEPH_PASS"),
		os.Getenv("PSQL_HOST")))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	//tx := db.MustBegin()

	cats := make([]Cat, 0)
	//var cat Cat

	err = db.Ping()
	if err != nil {
		return
	}
	err = db.Select(&cats, "SELECT * FROM cats")
	if err != nil {
		return
	}
	//rows, err := db.Query("SELECT * FROM cats")
	//for rows.Next() {
	//	err := rows.Scan(&cat.Name, &cat.Age, &cat.Color)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//	cats = append(cats, cat)
	//}

	fmt.Println(cats)
}
