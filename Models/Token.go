package models

type Token struct {
	Id          int    `json:"Id"`
	User_Id     int    `json:"User_Id"`
	Role        string `json:"role"`
	Email       string `json:"email" gorm:"Primarykey"`
	TokenString string `json:"token"`
}
