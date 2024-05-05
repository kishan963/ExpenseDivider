package database

import (
	"fmt"

	m "github.com/kishan963/Splitwise/Models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbSetup() error {
	fmt.Println("Db setup started")
	dsn := "postgresql://neondb_owner:9mY2EveIiQDu@ep-dawn-river-a54831g6.us-east-2.aws.neon.tech/neondb?sslmode=require"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db.AutoMigrate(&m.Group{}, &m.User{}, &m.Expense{}, &m.Balance{}, &m.Token{})
	if err != nil {
		return err
	}
	return nil

}

// func WriteToUserDb(user m.User) {
// 	Db.Create(&user)
// }
