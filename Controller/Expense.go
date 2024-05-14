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
				Group_Id:  data.GroupId,
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

func DeleteExpense(data m.Expense) {
	// d.Db.Where("Id = ?", data.Id).Delete(&data)
	var balance m.Balance
	d.Db.Where("Expenseid = ?", data.Id).Delete(&balance)
	var expense m.Expense
	if err := d.Db.First(&expense, data.Id).Error; err != nil {
		return // return error if expense is not found or any other error occurs
	}

	// Step 2: Remove association from user_expense table
	if err := d.Db.Model(&expense).Association("Users").Clear(); err != nil {
		return // return error if clearing association fails
	}

	// Step 3: Delete the expense
	if err := d.Db.Delete(&expense).Error; err != nil {
		return // return error if deletion fails
	}
}
