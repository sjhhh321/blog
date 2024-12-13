package models

type ArticleModel struct {
	Model
	Title        string    `gorm:"size:32" json:"title"`
	Abstract     string    `gorm:"size:256" json:"abstract"`
	Content      string    `json:"content"`
	CategoryID   uint      `json:"categoryID"`                                   //分类的id
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList"` //标签列表
	Cover        string    `gorm:"size:256" json:"cover"`
	UserID       uint      `json:"userId"`
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int       `json:"lookCount"`    //浏览数
	DiggCount    int       `json:"diggCount"`    //点赞数
	CommentCount int       `json:"commentCount"` //评论数
	OpenComment  bool      `json:"openComment"`  //是否开启评论
	Status       int8      `json:"status"`       //状态分为：草稿、审核中、已发布
}
