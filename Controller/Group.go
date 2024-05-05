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
