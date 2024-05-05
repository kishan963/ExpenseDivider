package controller

import (
	"fmt"

	d "github.com/kishan963/Splitwise/Database"
	m "github.com/kishan963/Splitwise/Models"
)

func Register(data m.User) error {
	fmt.Println("Registration started")

	// Unmarshal JSON data into RegistrationData struct

	// Process registration
	user := new(m.User)
	user.Id = data.Id
	user.Email = data.Email
	user.Phone = data.Phone
	user.Username = data.Username
	d.Db.Create(&user)
	fmt.Printf("Registration successful ", &user)
	return nil
}
