package models

type Balance struct {
	Id        int `gorm:"Primarykey"`
	Expenseid int
	By_user   int
	For_user  int
	Amount    float64
}
