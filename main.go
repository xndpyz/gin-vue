package main

import (
	"github.com/gin-gonic/gin"
	"lovelm/controller"
)


func main() {
	r := gin.Default()

	r.POST("/api/auth/register", controller.Register)
	panic(r.Run())
}

//func InitDB() *gorm.DB {
//	//driverName := "mysql"
//	//host := "localhost"
//	//port := "3306"
//	//database := ""
//
//}