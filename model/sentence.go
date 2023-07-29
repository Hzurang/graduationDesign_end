package model

import (
	"gorm.io/gorm"
	"time"
)

type Sentence struct {
	ID                  uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:每日一句表主键" json:"id"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	SentenceId          uint64         `gorm:"column:sentence_id;index;type:bigint(20) unsigned;comment:句子id，雪花算法生成" json:"sentence_id"`
	PublishAt           string         `gorm:"column:publish_at;type:varchar(16);comment:发布日期" json:"publish_at"`
	SentenceContent     string         `gorm:"column:sentence_content;type:varchar(255);comment:句子内容" json:"sentence_content"`
	SentenceNote        string         `gorm:"column:sentence_note;type:varchar(255);comment:句子译文" json:"sentence_note"`
	SentenceTranslation string         `gorm:"column:sentence_translation;type:varchar(255);comment:小编的话" json:"sentence_translation"`
	SentencePicture     string         `gorm:"column:sentence_picture;type:varchar(255);comment:句子图片链接" json:"sentence_picture"`
	SentenceAudioPath   string         `gorm:"column:sentence_audio_path;type:varchar(255);comment:句子语音链接" json:"sentence_audio_path"`
	SentenceCollectNum  uint64         `gorm:"column:sentence_collect_num;type:bigint(20) unsigned;not null;DEFAULT:0;comment:句子收藏数，默认为0" json:"sentence_collect_num"`
	DeleteIsOk          int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:句子是否删除，默认为0，1是删除" json:"delete_isok"`
}

type SentenceInfo struct {
	SentenceId          uint64 `json:"sentence_id"`
	PublishAt           string `json:"publish_at"`
	SentenceContent     string `json:"sentence_content"`
	SentenceNote        string `json:"sentence_note"`
	SentenceTranslation string `json:"sentence_translation"`
	SentencePicture     string `json:"sentence_picture"`
	SentenceAudioPath   string `json:"sentence_audio_path"`
	SentenceCollectNum  uint64 `json:"sentence_collect_num"`
}
