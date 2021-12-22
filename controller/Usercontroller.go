package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"lovelm/common"
	"lovelm/model"
	"lovelm/util"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	passwd := ctx.PostForm("passwd")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号应该为11位"})
		return
	}
	if len(passwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, passwd)
	// 判断手机号是否存在

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  passwd,
	}
	DB.Create(&newUser)

	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}
