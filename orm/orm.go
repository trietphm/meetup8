package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("postgres", "user=gopq password=123456 dbname=meetup sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect success")
	fmt.Println("---")

	// db.AutoMigrate(&User{})

	// Insert
	var user User
	user.Name = "Gopher"
	user.Age = 14
	err = db.Create(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("Insert success, User ID = %d \n", user.Id)
	fmt.Println("---")

	// Select
	var newUser User
	err = db.Model(User{Id: user.Id}).Find(&newUser).Error
	if err != nil {
		panic(err)
	}

	fmt.Printf("User: ID=%d, Name=%s, Age=%d\n", newUser.Id, newUser.Name, newUser.Age)

	// Update
	user.Age = 24
	user.Name = "John"
	err = db.Save(&user).Error
	if err != nil {
		panic(err)
	}
}
