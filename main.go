package main

import (
	"IOT_myblog/models" // 引入article_model

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite" // 引入sqlite
	"gorm.io/gorm"          // 引入gorm
)

func main() {
	r := gin.Default()

	//连数据库,此处把数据库命名为“yc”（原创）
	yc, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("可恶TT数据库连不上！")
	}

	//接口
	yc.AutoMigrate(&models.Article{})

	//接入article模型
	r.GET("/hello", func(a *gin.Context) { //一个小彩蛋，把常作为上下文简称的字母“c”改为了a^^
		a.JSON(200, gin.H{
			"message": "恭喜！数据库已连通！",
		})
	})

	r.Run(":8080")
}
