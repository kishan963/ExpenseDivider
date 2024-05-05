package jwttoken

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	d "github.com/kishan963/Splitwise/Database"
	m "github.com/kishan963/Splitwise/Models"
)

func IsAuthorized(w http.ResponseWriter, r *http.Request) (int, error) {

	// var data m.User
	// err := json.NewDecoder(r.Body).Decode(&data)

	if r.Header["Token"] == nil {
		var err error
		err = fmt.Errorf("Invalid Token: %v", err)
		return -1, err
	}

	var mySigningKey = []byte(secretkey)

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			var err error
			err = fmt.Errorf("Invalid Token: %v", err)
			return err, err
		}
		return mySigningKey, nil
	})
	var User_Token m.Token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		d.Db.Where("token_string = ?", r.Header["Token"][0]).Find(&User_Token)
		fmt.Println("Id ? Userid ? email ? token ?", User_Token.Id, User_Token.User_Id, User_Token.Email, r.Header["Token"][0])
		return User_Token.User_Id, nil
	}

	err = fmt.Errorf("Not authorized: %v", err)
	return -1, err
}
