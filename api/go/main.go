package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/boyan.bonev/intern-hackday/api/go/routes"
)

func main() {
	r := gin.Default()
	r.POST("/", routes.AnalyzeHandler)
	r.Run(":8080")
}
