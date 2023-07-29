package model

import (
	"gorm.io/gorm"
	"time"
)

type Essay struct {
	ID              uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:文章表主键" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	EssayId         uint64         `gorm:"column:essay_id;index;type:bigint(20) unsigned;comment:文章id，雪花算法生成" json:"essay_id"`
	EssayTitle      string         `gorm:"column:essay_title;type:varchar(255);comment:文章标题" json:"essay_title"`
	EssayAuthor     string         `gorm:"column:essay_author;type:varchar(16);comment:文章作者" json:"essay_author"`
	PublishAt       time.Time      `gorm:"column:publish_at;type:datetime(3);comment:发布日期" json:"publish_at"`
	EssayContent    string         `gorm:"column:essay_content;type:longtext;comment:文章内容" json:"essay_content"`
	EssayIsOk       int8           `gorm:"column:essay_isok;type:tinyint(1);not null;DEFAULT:0;comment:文章是否可评论，默认为0，1为可评论，0不可评论（预留字段）" json:"essay_isok"`
	EssayType       int8           `gorm:"column:essay_type;type:tinyint(1);comment:文章类型，0为英语小说，其他待定" json:"essay_type"`
	EssayCollectNum uint64         `gorm:"column:essay_collect_num;type:bigint(20) unsigned;not null;DEFAULT:0;comment:文章收藏数，默认为0" json:"essay_collect_num"`
	DeleteIsOk      int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:文章是否删除，默认为0，1是删除" json:"delete_isok"`
}

type EssayInfo struct {
	EssayId         uint64    `json:"essay_id"`
	EssayTitle      string    `json:"essay_title"`
	EssayAuthor     string    `json:"essay_author"`
	PublishAt       time.Time `json:"publish_at"`
	EssayContent    string    `json:"essay_content"`
	EssayIsOk       int8      `json:"essay_isok"`
	EssayType       int8      `json:"essay_type"`
	EssayCollectNum uint64    `json:"essay_collect_num"`
}

type EssayDetail struct {
	EssayId      string `json:"essay_id"`
	EssayTitle   string `json:"essay_title"`
	EssayAuthor  string `json:"essay_author"`
	PublishAt    string `json:"publish_at"`
	EssayContent string `json:"essay_content"`
	EssayType    int8   `json:"essay_type"`
	IsCollect    int8   `json:"is_collect"`
}

type EssayCollectInfo struct {
	EssayId         uint64    `json:"essay_id"`
	EssayTitle      string    `json:"essay_title"`
	EssayAuthor     string    `json:"essay_author"`
	PublishAt       time.Time `json:"publish_at"`
	EssayType       int8      `json:"essay_type"`
	EssayCollectNum uint64    `json:"essay_collect_num"`
}

type EssayList struct {
	EssayId         string `json:"essay_id"`
	EssayTitle      string `json:"essay_title"`
	EssayAuthor     string `json:"essay_author"`
	PublishAt       string `json:"publish_at"`
	EssayCollectNum string `json:"essay_collect_num"`
}
