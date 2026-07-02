package models

type ArticleModel struct {
	ID       int64
	Title    string     `gorm:"size: 20"`
	TagModel []TagModel `gorm:"many2many:article_tags;"`
}

type TagModel struct {
	ID          int64
	Title       string         `gorm:"size: 20"`
	ArticleList []ArticleModel `gorm:"many2many:article_tags;"`
}
