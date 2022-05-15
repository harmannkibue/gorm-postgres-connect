package main

import (
	"example/postgres-connect/pkg/books"
	"example/postgres-connect/pkg/common/config"
	"example/postgres-connect/pkg/common/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	configVariables, err := config.LoadConfig("./pkg/common/config/envs")

	if err != nil {
		log.Fatalln("Failed to load environment configs ", err)
	}

	port := configVariables.Port
	var db_host, db_user, password, db_name, db_port, ssl_mode, timezone = configVariables.DbHost, configVariables.DbUser, configVariables.Password, configVariables.DbName, configVariables.DbPort, configVariables.SslMode, configVariables.TimeZone

	db_init := db.Init("host=" + db_host + " user=" + db_user + " password=" + password + " dbname=" + db_name + " port=" + db_port + " sslmode=" + ssl_mode + " TimeZone=" + timezone)

	books.RegisterRoutes(router, db_init)

	log.Fatal(router.Run(port))

}
