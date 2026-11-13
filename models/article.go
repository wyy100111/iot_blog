package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model //内置函数

	Title   string `json:"title"`   // 标题
	Content string `json:"content"` // 文章内容
	Author  string `json:"author"`  // 作者
}
