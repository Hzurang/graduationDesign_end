package model

import (
	"gorm.io/gorm"
	"time"
)

type UserFeedback struct {
	ID              uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:用户反馈表主键" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	UserId          uint64         `gorm:"column:user_id;index;type:bigint(20) unsigned;comment:用户id，雪花算法生成" json:"user_id"`
	FeedbackContent string         `gorm:"column:feedback_content;type:longtext;not null;comment:反馈内容" json:"feedback_content"`
}
