package model

import (
	"gorm.io/gorm"
	"time"
)

type EssayWord struct {
	ID           uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:文章单词表主键" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	Word         string         `gorm:"column:word;type:varchar(32);comment:单词" json:"word"`
	WordNum      int            `gorm:"column:word_num;type:int;comment:单词序号" json:"word_num"`
	WordId       uint64         `gorm:"column:word_id;index;type:bigint(20) unsigned;comment:单词id，雪花算法生成" json:"word_id"`
	WordMusic    string         `gorm:"column:word_music;type:varchar(255);comment:单词发音链接" json:"word_music"`
	EssayId      uint64         `gorm:"column:essay_id;index;type:bigint(20) unsigned;comment:文章id" json:"essay_id"`
	WordMeaning  string         `gorm:"column:word_meaning;type:varchar(255);comment:单词意思" json:"word_meaning"`
	WordSentence string         `gorm:"column:word_sentence;type:longtext;comment:单词例句" json:"word_sentence"`
	DeleteIsOk   int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:单词是否删除，默认为0，1是删除" json:"delete_isok"`
}

type EssayWordInfo struct {
	WordId       uint64 `json:"word_id"`
	Word         string `json:"word"`
	WordNum      int    `json:"word_num"`
	WordMusic    string `json:"word_music"`
	EssayId      uint64 `json:"essay_id"`
	WordMeaning  string `json:"word_meaning"`
	WordSentence string `json:"word_sentence"`
}
