package model

import (
	"gorm.io/gorm"
	"time"
)

type Listen struct {
	ID               uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:听力表主键" json:"id"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	ListenId         uint64         `gorm:"column:listen_id;index;type:bigint(20) unsigned;comment:听力id，雪花算法生成" json:"listen_id"`
	ListenTitle      string         `gorm:"column:listen_title;type:varchar(255);comment:听力标题" json:"listen_title"`
	ListenSource     string         `gorm:"column:listen_source;type:varchar(16);comment:听力来源" json:"listen_source"`
	ListenEditor     string         `gorm:"column:listen_editor;type:varchar(16);comment:编辑" json:"listen_editor"`
	PublishAt        time.Time      `gorm:"column:publish_at;type:datetime(3);comment:发布日期" json:"publish_at"`
	ListenContent    string         `gorm:"column:listen_content;type:longtext;comment:听力内容" json:"listen_content"`
	ListenMediaPath  string         `gorm:"column:listen_media_path;type:varchar(255);comment:听力视频链接" json:"listen_media_path"`
	ListenMp3Path    string         `gorm:"column:listen_mp_3_path;type:varchar(255);comment:听力音频链接" json:"listen_mp_3_path"`
	ListenType       int8           `gorm:"column:listen_type;type:tinyint(1);comment:听力类型，0为热点资讯传送门，1为国外媒体资讯，2为英语听力入门，3为可可之声，4为品牌英语听力" json:"listen_type"`
	ListenSecondType string         `gorm:"column:listen_second_type;type:varchar(48);comment:听力第二级类型" json:"listen_second_type"`
	ListenCollectNum uint64         `gorm:"column:listen_collect_num;type:bigint(20) unsigned;not null;DEFAULT:0;comment:听力收藏数，默认为0" json:"listen_collect_num"`
	DeleteIsOk       int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:听力是否删除，默认为0，1是删除" json:"delete_isok"`
}

type ListenInfo struct {
	ListenId         uint64    `json:"listen_id"`
	ListenTitle      string    `json:"listen_title"`
	ListenSource     string    `json:"listen_source"`
	ListenEditor     string    `json:"listen_editor"`
	PublishAt        time.Time `json:"publish_at"`
	ListenContent    string    `json:"listen_content"`
	ListenMediaPath  string    `json:"listen_media_path"`
	ListenMp3Path    string    `json:"listen_mp_3_path"`
	ListenType       int8      `json:"listen_type"`
	ListenSecondType string    `json:"listen_second_type"`
	ListenCollectNum uint64    `json:"listen_collect_num"`
}

type ListenCacheInfo struct {
	ListenId         uint64    `json:"listen_id"`
	ListenTitle      string    `json:"listen_title"`
	ListenSource     string    `json:"listen_source"`
	ListenEditor     string    `json:"listen_editor"`
	PublishAt        time.Time `json:"publish_at"`
	ListenSecondType string    `json:"listen_second_type"`
	ListenCollectNum uint64    `json:"listen_collect_num"`
}
