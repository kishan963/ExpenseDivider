package dto

type ResponseDto struct {
	Id            int `gorm:"Primarykey"`
	Expenseid     int
	By_user       int
	By_user_name  string
	For_user      int
	For_user_name string
	Group_Id      int
	Amount        float64
}
