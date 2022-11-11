package controller

import (
	"bookmanager/dao/mysql"
	"bookmanager/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// 注册
func RegisterHandler(c *gin.Context) {
	p := &model.User{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//账号密码落库
	if tx := mysql.DB.Create(p); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": tx.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"data": p.Username,
	})
}

// 登录
func LoginHandler(c *gin.Context) {
	p := &model.User{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//校验账号密码
	rows := mysql.DB.Model(p).First(&p).Row()
	if rows == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "账号密码错误",
		})
		return
	}
	//验证通过生成token
	token := uuid.New().String()
	if tx := mysql.DB.Model(p).Update("token", token); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})
}
