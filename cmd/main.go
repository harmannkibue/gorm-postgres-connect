package main

import (
	"example/postgres-connect/pkg/books"
	"example/postgres-connect/pkg/common/config"
	"example/postgres-connect/pkg/common/db"
	"example/postgres-connect/pkg/common/utils"
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

	dbInit := db.Init(utils.BuildPostgresParams())

	books.RegisterRoutes(router, dbInit)

	log.Fatal(router.Run(port))

}
