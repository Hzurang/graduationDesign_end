package model

import (
	"gorm.io/gorm"
	"time"
)

type ListenWord struct {
	ID           uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:听力单词表主键" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	WordId       uint64         `gorm:"column:word_id;index;type:bigint(20) unsigned;comment:单词id，雪花算法生成" json:"word_id"`
	Word         string         `gorm:"column:word;type:varchar(32);comment:单词" json:"word"`
	WordPhonetic string         `gorm:"column:word_phonetic;type:varchar(32);comment:单词音标" json:"word_phonetic"`
	WordMeaning  string         `gorm:"column:word_meaning;type:varchar(255);comment:单词意思" json:"word_meaning"`
	WordMusic    string         `gorm:"column:word_music;type:varchar(255);comment:单词发音链接" json:"word_music"`
	WordNum      int            `gorm:"column:word_num;type:int;comment:单词序号" json:"word_num"`
	ListenId     uint64         `gorm:"column:listen_id;index;type:bigint(20) unsigned;comment:听力id" json:"listen_id"`
	DeleteIsOk   int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:单词是否删除，默认为0，1是删除" json:"delete_isok"`
}

type ListenWordInfo struct {
	WordId       uint64 `json:"word_id"`
	Word         string `json:"word"`
	WordPhonetic string `json:"word_phonetic"`
	WordMeaning  string `json:"word_meaning"`
	WordMusic    string `json:"word_music"`
	WordNum      int    `json:"word_num"`
	ListenId     uint64 `json:"listen_id"`
}
