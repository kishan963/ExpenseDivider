package models

type Group struct {
	Id       int `gorm:"Primarykey"`
	Name     string
	Expenses []Expense `gorm:"foreignKey:GroupId"`
	Users    []User    `gorm:"many2many:user_group;"`
}
