package logic

import (
	"code.oldboy.com/studygolang/day12/blogger/db"
	"code.oldboy.com/studygolang/day12/blogger/models"
)

// 文章分类

func GetCategoryList() (categoryList []*models.Category, err error) {
	return db.GetAllCategoryList()
}
