package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	if isTelephoneExit(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已存在"})
		return
	}

	// 创建用户
	hashPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "密码加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPasswd),
	}
	DB.Create(&newUser)

	ctx.JSON(200, gin.H{
		"message": "注册成功",
		"code":    200,
	})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	telephone := ctx.PostForm("telephone")
	passwd := ctx.PostForm("passwd")

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号未注册"})
		return
	}

	// 判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 添加token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "系统错误"})
		log.Println("token generate error:", err)
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"message": "登陆成功",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}

func isTelephoneExit(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
