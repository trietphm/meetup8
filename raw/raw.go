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

	// Insert
	var userId int
	err = db.QueryRow(`INSERT INTO users(name, age)	VALUES('Gopher', 14) RETURNING id`).Scan(&userId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Insert success, User ID = %d \n", userId)
	fmt.Println("---")

	// Select
	row, err := db.Query("SELECT id, name, age FROM users WHERE id = $1", userId)
	if err != nil {
		panic(err)
	}

	var id int
	var name string
	var age int

	defer row.Close()
	for row.Next() {
		err = row.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
	}
	if row.Err() != nil {
		panic(err)
	}
	fmt.Printf("User: ID=%d, Name=%s, Age=%d\n", id, name, age)

	// Update
	_, err = db.Exec("UPDATE users SET name = 'John' WHERE id = $1", userId)
	if err != nil {
		panic(err)
	}
}
