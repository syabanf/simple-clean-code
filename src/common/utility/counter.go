package utility

type Count struct {
	ArticleID string `gorm:"article_id"`
	Count     int64  `gorm:"count"`
}
