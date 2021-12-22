package main

import (
	"github.com/gin-gonic/gin"
	"lovelm/common"
	"lovelm/controller"
	"lovelm/middleware"
)

func main() {
	r := gin.Default()
	db := common.InitDB()
	defer db.Close()
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	panic(r.Run())
}
