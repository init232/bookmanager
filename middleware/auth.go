package middleware

import (
	"bookmanager/dao/mysql"
	"bookmanager/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		//验证token是否正取
		var u *model.User
		row := mysql.DB.Where("token = ?", token).First(&u).RowsAffected
		if row != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "验证失败",
			})
			c.Abort()
		}
		c.Set("UserId", u.ID)
	}
}
