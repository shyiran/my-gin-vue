package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"shyiran/my-gin-vue/common"
	"shyiran/my-gin-vue/dao"
	"shyiran/my-gin-vue/db"
	"shyiran/my-gin-vue/dto"
	"shyiran/my-gin-vue/model"
	"shyiran/my-gin-vue/response"
	"shyiran/my-gin-vue/util"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	if len(phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户手机号码必须11位数字!")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户密码不能少于6位!")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if dao.IsPhoneExist(phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "User exisit!")
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "Hased password error!")
		return
	}
	user := model.User{
		Name:     name,
		Password: string(hasedPassword),
		Phone:    phone,
	}
	db.DB.Create(&user)
	response.Succces(ctx, nil, "Register success!")
}

func Login(c *gin.Context) {
	phone := c.PostForm("phone")       // 获取参数
	password := c.PostForm("password") //获取密码
	if phone == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Phone num not null!"})
		return
	}
	if password == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password not null!"})
		return
	}
	if len(phone) != 11 { // 数据验证
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Phone num must 11 digits!"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位!"})
		return
	}
	var user model.User // 判断手机号是否存在
	db.DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // 判断密码是否正确
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password err!"})
	}
	token, err := common.ReleaseToken(user) // 发放token
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "System err!"})
		log.Printf("token generate error:%v", err)
		return
	}
	response.Succces(c, gin.H{"token": token}, "Login success!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
