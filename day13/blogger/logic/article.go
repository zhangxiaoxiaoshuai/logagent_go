package logic

import (
	"code.oldboy.com/studygolang/day12/blogger/db"
	"code.oldboy.com/studygolang/day12/blogger/models"
)

// 业务逻辑层

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*models.ArticleRecord, err error) {
	// 取文章数据
	articleList, err := db.GetArticleInfo(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) == 0 {
		return
	}

	// 获取文章分类信息
	categoryList, err := db.GetAllCategoryList()
	if err != nil {
		return
	}
	// 聚合数据:把文章和它对应的文章分类聚合成models.ArticleRecord
	articleRecordList = make([]*models.ArticleRecord, 0, len(articleList))
	for _, a := range articleList {
		for _, c := range categoryList {
			if a.CategoryId == c.CategoryId {
				tmpC := &models.ArticleRecord{
					ArticleInfo: *a,
					Category:    *c,
				}
				articleRecordList = append(articleRecordList, tmpC)
				break
			}
		}
	}
	return
}
