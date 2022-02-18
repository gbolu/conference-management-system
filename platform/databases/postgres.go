package database

import (
	"fmt"

	"github.com/gbolu/conference-management-system/pkg/configs"
	database_utils "github.com/gbolu/conference-management-system/platform"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgreSQLConnection()  {
	// Define database connection for PostgreSQL.
	var err error
	dsn := configs.GetEnvVariable("DB_SERVER_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if(err != nil) {
		panic(fmt.Errorf("error, not connected to database, %w", err))
	}

	fmt.Println("Database connection established successfully.")
	
	database_utils.SetDatabase(db)
}
