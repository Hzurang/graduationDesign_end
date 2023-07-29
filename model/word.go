package model

import (
	"gorm.io/gorm"
	"time"
)

type Word struct {
	ID               uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:单词表主键" json:"id"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	WordId           uint64         `gorm:"column:word_id;index;type:bigint(20) unsigned;comment:单词id，雪花算法生成" json:"word_id"`
	Word             string         `gorm:"column:word;type:varchar(48);comment:单词" json:"word"`
	PhoneticTransEng string         `gorm:"column:phonetic_trans_eng;type:varchar(32);comment:单词音标（英）" json:"phonetic_trans_eng"`
	PhoneticTransAme string         `gorm:"column:phonetic_trans_ame;type:varchar(32);comment:单词音标（美）" json:"phonetic_trans_ame"`
	WordMeaning      string         `gorm:"column:word_meaning;type:varchar(144);comment:单词意思" json:"word_meaning"`
	MnemonicAid      string         `gorm:"column:mnemonic_aid;type:longtext;comment:单词助记" json:"mnemonic_aid"`
	ChiEtymology     string         `gorm:"column:chi_etymology;type:longtext;comment:单词中文词源" json:"chi_etymology"`
	SentenceEng1     string         `gorm:"column:sentence_eng_1;type:longtext;comment:单词例句英文1" json:"sentence_eng_1"`
	SentenceChi1     string         `gorm:"column:sentence_chi_1;type:longtext;comment:单词例句中文1" json:"sentence_chi_1"`
	SentenceEng2     string         `gorm:"column:sentence_eng_2;type:longtext;comment:单词例句英文2" json:"sentence_eng_2"`
	SentenceChi2     string         `gorm:"column:sentence_chi_2;type:longtext;comment:单词例句中文2" json:"sentence_chi_2"`
	SentenceEng3     string         `gorm:"column:sentence_eng_3;type:longtext;comment:单词例句英文3" json:"sentence_eng_3"`
	SentenceChi3     string         `gorm:"column:sentence_chi_3;type:longtext;comment:单词例句中文3" json:"sentence_chi_3"`
	WordType         int8           `gorm:"column:word_type;type:tinyint(1);not null;DEFAULT:0;comment:单词类型（所属词书），默认为0，1为四级词汇，2为六级词汇，3为英专四级，4为英专八级，5为考研词汇，6为GRE词汇，7为托福词汇，8为雅思词汇" json:"word_type"`
	DeleteIsOk       int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:单词是否删除，默认为0，1是删除" json:"delete_isok"`
}

type WordCollectInfo struct {
	WordId      uint64 `json:"word_id"`
	Word        string `json:"word"`
	WordMeaning string `json:"word_meaning"`
	WordType    int8   `json:"word_type"`
}

type WordInfo struct {
	WordId           uint64 `json:"word_id"`
	Word             string `json:"word"`
	PhoneticTransEng string `json:"phonetic_trans_eng"`
	PhoneticTransAme string `json:"phonetic_trans_ame"`
	WordMeaning      string `json:"word_meaning"`
	MnemonicAid      string `json:"mnemonic_aid"`
	ChiEtymology     string `json:"chi_etymology"`
	SentenceEng1     string `json:"sentence_eng_1"`
	SentenceChi1     string `json:"sentence_chi_1"`
	SentenceEng2     string `json:"sentence_eng_2"`
	SentenceChi2     string `json:"sentence_chi_2"`
	SentenceEng3     string `json:"sentence_eng_3"`
	SentenceChi3     string `json:"sentence_chi_3"`
	WordType         int8   `json:"word_type"`
}
