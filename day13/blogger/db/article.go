package db

import (
	"fmt"

	"code.oldboy.com/studygolang/day12/blogger/models"
)

// article action

// GetArticleInfo 获取文章信息
func GetArticleInfo(pageNum, pageSize int) (articleInfoList []*models.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid param,page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlStr := `
		select
			id, summary, title, view_count, create_time, comment_count, username, category_id
		from
			article
		where
			status = 1
		order by create_time desc
		limit ?,?
	`
	err = DB.Select(&articleInfoList, sqlStr, pageNum, pageSize)
	return
}
