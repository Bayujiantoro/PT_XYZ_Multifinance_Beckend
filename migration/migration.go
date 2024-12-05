package migration

import (
	"fmt"
	"pt-xyz-multifinance/connection"
	"pt-xyz-multifinance/entity/model"
)

func RunMigration() {
	err := connection.DB.AutoMigrate(
		&model.Users{},
		&model.Tenor{},
		&model.Transaction{},
		&model.Pinjaman{},
		&model.Product{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}