package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbName := os.Getenv("DB_Name")
	dbPassword := os.Getenv("DB_Password")
	conn := fmt.Sprintf("raihan:%v@tcp(localhost:3306)/%v", dbPassword, dbName)
	fmt.Println(conn)
	DB, err := sql.Open("mysql", conn)
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
