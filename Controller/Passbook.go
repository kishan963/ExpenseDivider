package controller

import (
	"fmt"
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
	fmt.Println(balances)
	for j, v := range balances {
		if v.For_user == Id && v.Amount > 0 {
			for i, u := range balances {
				if u.For_user == Id && u.By_user == v.By_user && i > j {
					balances[j].Amount += u.Amount
					balances[i].Amount = 0
				}
			}

		}
	}

	for j, v := range balances {
		if v.By_user == Id && v.Amount > 0 {
			for i, u := range balances {
				if u.By_user == Id && u.For_user == v.For_user && i > j {
					balances[j].Amount += u.Amount
					balances[i].Amount = 0
				}
			}

		}
	}

	var updated_balance []m.Balance

	for _, v := range balances {
		if v.Amount != 0 && v.By_user != v.For_user {
			updated_balance = append(updated_balance, v)
		}
	}
	return updated_balance
}
