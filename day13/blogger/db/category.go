package db

import "code.oldboy.com/studygolang/day12/blogger/models"

// category actions

// GetAllCategoryList 查询所有的分类信息
func GetAllCategoryList() (categoryList []*models.Category, err error) {
	sqlStr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	if err != nil {
		return
	}
	return
}
