package main

import (
	"financial-management/config"
	"financial-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.DB()

	routes.Api(r, db)

	r.Run()
}
