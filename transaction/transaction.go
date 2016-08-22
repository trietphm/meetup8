package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=gopq password=123456 dbname=meetup sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect success")
	fmt.Println("---")

	Transaction(db)
}

func Transaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Insert
	var userId int
	err = tx.QueryRow(`INSERT INTO users(name, age)	VALUES('Gopher', 14) RETURNING id`).Scan(&userId)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	fmt.Printf("Insert success, User ID = %d \n", userId)
	fmt.Println("---")

	// Update
	_, err = tx.Exec("UPDATE users SET name = 'John' WHERE id = $1", userId)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	// Select
	row, err := tx.Query("SELECT id, name, age FROM users WHERE id = $1", userId)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	var id int
	var name string
	var age int

	defer row.Close()
	for row.Next() {
		err = row.Scan(&id, &name, &age)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
	if row.Err() != nil {
		tx.Rollback()
		panic(err)
	}
	fmt.Printf("User: ID=%d, Name=%s, Age=%d\n", id, name, age)

	tx.Commit()
}
