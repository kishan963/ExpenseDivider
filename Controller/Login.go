package controller

import (
	"encoding/json"
	"fmt"

	d "github.com/kishan963/Splitwise/Database"
	j "github.com/kishan963/Splitwise/JwtToken"
	m "github.com/kishan963/Splitwise/Models"
)

func Login(data m.User) (json.Token, error) {
	fmt.Println("Login started")
	var users m.User

	if err := d.Db.Where("phone = ? AND email = ?", data.Phone, data.Email).Find(&users).Error; err != nil || users.Id == 0 {
		err = fmt.Errorf("Invalid email and phone: %v", err)
		return nil, err
	}

	validToken, err := j.GenerateJWT(data.Email, "User", users.Id)
	if err != nil {
		err = fmt.Errorf("failed to generate jwt token: %v", err)
		return nil, err
	}

	var token m.Token
	var token1 m.Token
	token.Email = data.Email
	token.Role = "User"
	token.User_Id = users.Id
	token.TokenString = validToken
	d.Db.Where("user_id = ?", users.Id).Delete(&token1)
	d.Db.Create(&token)
	return token, nil
}
