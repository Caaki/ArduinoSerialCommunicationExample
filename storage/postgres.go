package storage

import (
	"SerialArduinoCommunication/configuration"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Repository ...
// TODO: Proveriti da li se uklanja ili ostaje!
type Repository struct {
	DB *gorm.DB
}

var Repo Repository

func CreteConnection() error {
	config := configuration.GlobalConfiguration

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.PostgresQL.Host,
		config.PostgresQL.Port,
		config.PostgresQL.User,
		config.PostgresQL.Password,
		config.PostgresQL.DBName,
		config.PostgresQL.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Repo.DB = db
	return nil
}
