package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	dbUrl:= "postgres://raihan:secret@localhost/golang?sslmode=disable"
	db, err:= sql.Open("postgres", dbUrl)
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	err= db.Ping()
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("connected database")
}
