package database

import (
	"fmt"

	"github.com/sagarmaheshwary/learning-golang/articles-app-with-jwt/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dbConfig := config.GetDBConfig()

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.Port,
		dbConfig.SSLMode,
		dbConfig.TimeZone,
	)

	var err error

	//assigning connection to "DB" variable that's declared in ./.database.go
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.Raw("SELECT 1+1;")

	fmt.Println("Database Connected.")
}
