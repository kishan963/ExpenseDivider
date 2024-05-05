package models

type User struct {
	Id       int `gorm:"Primarykey"`
	Username string
	Email    string
	Phone    string
	Groups   []Group   `gorm:"many2many:user_group;"`
	Expenses []Expense `gorm:"many2many:user_expense;"`
}
