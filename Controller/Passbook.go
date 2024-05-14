package controller

import (
	"fmt"
	"log"

	"github.com/jinzhu/copier"
	d "github.com/kishan963/Splitwise/Database"
	dto "github.com/kishan963/Splitwise/Dto"
	m "github.com/kishan963/Splitwise/Models"
)

func GetUserBalance(data m.Expense) []dto.ResponseDto {
	var balances []m.Balance
	var responseDto []dto.ResponseDto
	// Retrieve all data from the users table
	if err := d.Db.Where("group_id = ? AND (for_user = ? OR by_user = ?)", data.GroupId, data.Id, data.Id).Find(&balances).Error; err != nil {
		log.Fatal(err)
	}
	balances = SortUserBalance(balances, data.Id)
	copier.Copy(&responseDto, &balances)
	responseDto = SetUserName(responseDto)
	return responseDto

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

func SetUserName(responseDto []dto.ResponseDto) []dto.ResponseDto {
	var user []m.User
	user = GetAllUser()

	userMap := make(map[int]string)
	for _, u := range user {
		userMap[u.Id] = u.Username
	}

	for i, v := range responseDto {
		responseDto[i].By_user_name = userMap[v.By_user]
		responseDto[i].For_user_name = userMap[v.For_user]
	}
	return responseDto
}
