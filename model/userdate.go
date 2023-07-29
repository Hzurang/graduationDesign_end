package model

import (
	"gorm.io/gorm"
	"time"
)

type UserDate struct {
	ID               uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:用户打卡情况表主键" json:"id"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	UserId           uint64         `gorm:"column:user_id;index;type:bigint(20) unsigned;comment:用户id，雪花算法生成，无符号" json:"user_id"`
	Date             time.Time      `gorm:"column:date;type:datetime(3);comment:时间" json:"date"`
	WordLearnNumber  int            `gorm:"column:word_learn_number;type:int;comment:在这一天新学多少单词" json:"word_learn_number"`
	WordReviewNumber int            `gorm:"column:word_review_number;type:int;comment:在这一天复习多少单词" json:"word_review_number"`
	Remark           string         `gorm:"column:remark;type:longtext;comment:在这一天的心情感悟" json:"remark"`
	DeleteIsOk       int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:用户信息是否删除，默认为0，1是删除" json:"delete_isok"`
}
