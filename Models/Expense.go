package models

type Expense struct {
	Id          int `gorm:"Primarykey"`
	Description string
	Amount      float64
	GroupId     int
	PaidBy_User int
	Users       []User `gorm:"many2many:user_expense;"`
}
