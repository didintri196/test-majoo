package main

import (
	"log"
	"majoo-test/middleware"
	"majoo-test/util"

	"github.com/gin-gonic/gin"
)

// @title API MINI-POS
// @version 1.0
// @description Api untuk service MINI-POS
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	gin.SetMode(gin.DebugMode)
	log.Print("Starting the majoo-test app")
	middleware.Middleware()
}
func init() {
	util.ConnectionDB()
}
