package controller

import (
	"bookmanager/dao/mysql"
	"bookmanager/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 新增
func CreateBookHandler(c *gin.Context) {
	p := &model.Book{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//数据库落库
	tx := mysql.DB.Create(&p)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建书籍成功",
	})
}

// 查看列表
func GetBookListHandler(c *gin.Context) {
	books := []*model.Book{}
	tx := mysql.DB.Preload("Users").Find(&books)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询列表成功",
		"data": books,
	})
}

// 查看书籍详情
func GetBookDetailHandler(c *gin.Context) {
	idstr := c.Param("id")
	//string转换成int
	bookID, _ := strconv.ParseInt(idstr, 10, 64)
	book := &model.Book{
		ID: bookID,
	}
	//数据库查询
	tx := mysql.DB.Preload("Users").Find(&book)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "查询书籍失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询书籍明细成功",
		"data": book,
	})
}

// 更新书籍
func UpdateBookHandler(c *gin.Context) {
	p := &model.Book{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "更新书籍失败",
		})
		return
	}
	//数据库更新
	oldBook := &model.Book{ID: p.ID}
	var newBook *model.Book
	if p.Name != "" {
		newBook.Name = p.Name
	}
	if p.Desc != "" {
		newBook.Desc = p.Desc
	}
	tx := mysql.DB.Model(&oldBook).Updates(&newBook)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新书籍成功",
	})
}

// 删除
func DeleteBookHandler(c *gin.Context) {
	idstr := c.Param("id")
	//str转换成int64
	bookID, _ := strconv.ParseInt(idstr, 10, 64)
	//数据库删除,删除book时同时删除第三张表中的对应关系
	if tx := mysql.DB.Select("Users").Delete(&model.Book{ID: bookID}); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "删除书籍成功",
		"data": bookID,
	})
}
