package models

import "time"

// ArticleInfo 存放基本的文章数据
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// ArticleDetail 包含文章基本信息、文章内容、文章分类信息的结构体
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

// ArticleRecord 包含文章基本信息和文章分类信息的结构体
type ArticleRecord struct {
	ArticleInfo
	Category
}
