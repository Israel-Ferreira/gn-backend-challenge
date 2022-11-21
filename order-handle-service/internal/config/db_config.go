package config

import (
	"fmt"
	"log"

	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func DbConfiguration() {
	dsn := createDbConnectionString(Envs.DbConfigEnv)

	conn, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("Error on Connect DB: %v \n", err)
	}

	DbConn = conn

	DbConn.AutoMigrate(&models.Order{}, &models.Address{})
}

func createDbConnectionString(dbConfig DbConfigEnv) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
		dbConfig.DbPort,
	)
}
