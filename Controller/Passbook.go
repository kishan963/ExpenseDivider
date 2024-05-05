package controller

import (
	"log"

	d "github.com/kishan963/Splitwise/Database"
	m "github.com/kishan963/Splitwise/Models"
)

func GetUserBalance(Id int) []m.Balance {
	var balances []m.Balance

	// Retrieve all data from the users table
	if err := d.Db.Where("for_user = ? OR 	by_user = ?", Id, Id).Find(&balances).Error; err != nil {
		log.Fatal(err)
	}
	return SortUserBalance(balances, Id)
}

func SortUserBalance(balances []m.Balance, Id int) []m.Balance {

	for j, v := range balances {
		if v.For_user == Id {
			for i, u := range balances {
				if u.For_user == v.By_user {
					balances[i].Amount -= v.Amount
					balances[j].Amount = 0
				}
			}

		}
	}

	var updated_balance []m.Balance

	for _, v := range balances {
		if v.Amount != 0 {
			updated_balance = append(updated_balance, v)
		}
	}
	return updated_balance
}
