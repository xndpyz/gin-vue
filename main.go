package main

import (
	"github.com/gin-gonic/gin"
	"lovelm/common"
	"lovelm/controller"
)

func main() {
	r := gin.Default()
	db := common.InitDB()
	defer db.Close()
	r.POST("/api/auth/register", controller.Register)
	panic(r.Run())
}
