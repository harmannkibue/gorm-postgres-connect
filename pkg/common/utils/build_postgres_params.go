package utils

import (
	"example/postgres-connect/pkg/common/config"
	"log"
)

// BuildPostgresParams returns configs in form; host="" user="" password="" dbname="" port="" sslmode="" TimeZone=""
func BuildPostgresParams() string {
	configVariables, err := config.LoadConfig("./pkg/common/config/envs")

	if err != nil {
		log.Fatalln("Failed to load environment variable configurations")
	}

	// keys to be passed to postgres  for initialization
	keys := []string{"host=", "user=", "password=", "dbname=", "port=", "sslmode=", "TimeZone="}
	// Load string with environment variables for postgres
	values := []string{configVariables.DbHost, configVariables.DbUser, configVariables.Password, configVariables.DbName, configVariables.DbPort, configVariables.SslMode, configVariables.TimeZone}
	var returnString string

	for i, ky := range keys {
		returnString += ky + values[i] + " "
	}

	return returnString
}
