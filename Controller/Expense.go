package controller

import (
	"fmt"

	d "github.com/kishan963/Splitwise/Database"
	m "github.com/kishan963/Splitwise/Models"
)

func AddExpense(data m.Expense) {

	d.Db.Create(&data)
	UpdateBalance(data)

}

func UpdateBalance(data m.Expense) {
	fmt.Println("Inside UpdateBalance ")
	for _, v := range data.Users {
		var balance m.Balance
		if result := d.Db.Where("by_user = ? AND for_user = ?", data.PaidBy_User, v.Id).Find(&balance); result.Error == nil {
			balance = m.Balance{
				Expenseid: data.Id,
				By_user:   data.PaidBy_User,
				For_user:  v.Id,
				Amount:    data.Amount / float64(len(data.Users)),
			}
			fmt.Println("Updatebalance ", result)
		} else {

			err := d.Db.Delete(balance)
			if err.Error != nil {
				fmt.Println(err.Error)
			}
			balance.Amount = balance.Amount + data.Amount/float64(len(data.Users))
		}

		err := d.Db.Create(&balance)
		if err.Error != nil {
			fmt.Println(err.Error)
		}
	}

}
