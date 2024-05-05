package controller

import (
	"fmt"
	"log"

	d "github.com/kishan963/Splitwise/Database"
	m "github.com/kishan963/Splitwise/Models"
)

func GetAllUser() []m.User {
	var users []m.User
	fmt.Println("Inside getAlluser")
	// Retrieve all data from the users table
	if err := d.Db.Preload("Groups").Preload("Expenses").Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("users println ", users)
	return users
}

func GetSingleUser() []m.User {
	var users []m.User

	// Retrieve all data from the users table
	if err := d.Db.Preload("Groups").Preload("Expenses").Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	return users
}

func DeleteExpense() {
	fmt.Println("Registation started")
}
