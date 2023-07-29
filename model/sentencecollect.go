package model

import (
	"gorm.io/gorm"
	"time"
)

type SentenceCollect struct {
	ID         uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:句子收藏表主键" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	SentenceId uint64         `gorm:"column:sentence_id;index;type:bigint(20) unsigned;comment:句子id，雪花算法生成" json:"sentence_id"`
	UserId     uint64         `gorm:"column:user_id;index;type:bigint(20) unsigned;comment:用户id，雪花算法生成" json:"user_id"`
	DeleteIsOk int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:收藏记录是否删除，默认为0，1是删除" json:"delete_isok"`
}

type SentenceCollectInfo struct {
	SentenceId         uint64 `json:"sentence_id"`
	PublishAt          string `json:"publish_at"`
	SentenceContent    string `json:"sentence_content"`
	SentenceCollectNum uint64 `json:"sentence_collect_num"`
}
