package controller

import (
	"fmt"

	d "github.com/kishan963/Splitwise/Database"
	m "github.com/kishan963/Splitwise/Models"
)

var GroupList []m.Group

func CreateGroup(data m.Group) error {
	d.Db.Create(&data)
	return nil
}

func GetGroup(data m.Group) (m.Group, error) {
	var group m.Group
	d.Db.Preload("Expenses").Preload("Users").Find(&group, data.Id)
	return group, nil
}

func GetUserGroups(data m.User) ([]m.Group, error) {
	fmt.Println("GetUserGroupp")
	var user m.User
	var groups []m.Group
	d.Db.Preload("Groups").Find(&user, data.Id)
	fmt.Println(user)
	groups = append(groups, user.Groups...)
	return groups, nil
}

func DeleteGroup(data m.Group) {
	// fmt.Println("Group Id Delete ", data)
	// d.Db.Where("Id = ?", data.Id).Delete(&data)

	var group m.Group
	if err := d.Db.Preload("Expenses").First(&group, data.Id).Error; err != nil {
		return // return error if expense is not found or any other error occurs
	}

	for _, v := range group.Expenses {
		DeleteExpense(v)
	}

	// Step 2: Remove association from user_expense table
	if err := d.Db.Model(&group).Association("Users").Clear(); err != nil {
		return // return error if clearing association fails
	}

	// Step 3: Delete the expense
	if err := d.Db.Delete(&group).Error; err != nil {
		return // return error if deletion fails
	}
}
