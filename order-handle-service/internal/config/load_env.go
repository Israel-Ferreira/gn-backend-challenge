package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type RabbitMqEnv struct {
	RabbitBrokerUrl string
}

type DbConfigEnv struct {
	DbPort     int64
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
}

type EnvConfig struct {
	RabbitMqEnv
	DbConfigEnv
}

var Envs = &EnvConfig{}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found...")
	}

	Envs.RabbitBrokerUrl = os.Getenv("RABBIT_MQ_BROKER_URL")


	Envs.DbName = os.Getenv("DB_NAME")
    Envs.DbHost = os.Getenv("DB_HOST")

    Envs.DbUser = os.Getenv("DB_USER")
    Envs.DbPassword = os.Getenv("DB_PASS")

	port := "5432"

	if envDbPort := os.Getenv("DB_PORT"); envDbPort != "" {
		port = fmt.Sprintf("%s", envDbPort)
	}

	portNum, err := strconv.ParseInt(port, 10, 64)

	if err != nil {
		log.Fatalf("Error on Convert DbPort In Int: %v", err)
	}

	Envs.DbPort = portNum

}
