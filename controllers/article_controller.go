package controllers

import (
	"IOT_myblog/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建接收JSON数据的结构体
type Create_article_input struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// 处理发文章请求
func Create_article(yc *gorm.DB) gin.HandlerFunc {
	return func(a *gin.Context) {
		var input Create_article_input

		//造新变量，用来塞数据
		article := models.Article{
			Title:   input.Title,
			Content: input.Content,
			Author:  input.Author,
		}

		err := a.ShouldBindJSON(&input) //查阅了gin官网的绑定函数
		if err != nil {
			a.JSON(400, gin.H{ //查阅资料可知，400代表输入的参数错误
				"code":    400, //返回前端的状态码
				"message": "参数错误",
				"error":   err.Error(),
			})
			return
		}

		//呼叫yc把实例存进数据库
		result := yc.Create(&article)
		var err_result error
		err_result = result.Error //gorm塞错误信息的地方
		if err_result != nil {
			a.JSON(500, gin.H{
				"code":    500,
				"message": "警告，数据库存不进去！",
			})
			return
		}

		//存成功了！
		a.JSON(200, gin.H{
			"code":    0,
			"message": "发布成功！",
			"data":    article,
		})
	}
}
