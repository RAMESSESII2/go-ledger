package repositories

import (
	"fmt"

	"github.com/RAMESSESII2/go-ledger/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitialMigration(DNS string) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	}
	fmt.Println("Connected to Databse")
	DB.AutoMigrate(&models.Transaction{})
}
