package main

import (
	// "goginapi/interfaces/api/http/customer_route"
	"goginapi/interfaces/api/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoute1Routes(r)
	// router.GET("/", getDetail)
	r.Run(":8001") // default 8080
}
