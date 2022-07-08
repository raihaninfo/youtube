package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB, err := sql.Open("mysql", "raihan:Raihan@@901@tcp(localhost:3306)/session19")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		log.Println(err.Error())
	}

	// insert data

	// model.DB.Close()
	// result, err := DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)",
	// 	"Rezaul", "rezaul@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(id)

	// Select all data
	// result, err:= DB.Query("SELECT * FROM users")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer result.Close()
	// for result.Next() {
	// 	var id int
	// 	var name string
	// 	var email string
	// 	err := result.Scan(&id, &name, &email)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	log.Println(id, name, email)
	// }

	// delete data

	result, err := DB.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffected)
}
