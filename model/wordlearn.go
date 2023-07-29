package model

import (
	"gorm.io/gorm"
	"time"
)

// 深度掌握次数
/*
 * 前提：掌握程度已达到10
 * 当深度次数为0时，记下次复习时间=上次已掌握时间+4天，若及时复习，更新上次已掌握时间
 * 当深度次数为1时，记下次复习时间=上次已掌握时间+3天，若及时复习，更新上次已掌握时间
 * 当深度次数为2时，记下次复习时间=上次已掌握时间+8天，若及时复习，更新上次已掌握时间
 * 当深度次数为3时，记已经完全掌握
 *
 * 检测哪些单词未及时深度复习：
 * 首先单词必须掌握程度=10，其次单词上次掌握的时间与现在的时间进行对比
 * （1）要是深度次数为0，且两者时间之差为大于4天，说明未深度复习
 * （2）要是深度次数为1，且两者时间之差为大于3天，说明未深度复习
 * （3）要是深度次数为2，且两者时间之差为大于8天，说明未深度复习
 * （4）若未及时深度复习，一律将其单词掌握程度-2（10→8）
 *
 * */

type WordLearn struct {
	ID              uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:用户单词学习情况表主键" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	WordId          uint64         `gorm:"column:word_id;index;type:bigint(20) unsigned;comment:单词id，雪花算法生成" json:"word_id"`
	UserId          uint64         `gorm:"column:user_id;index;type:bigint(20) unsigned;comment:用户id，雪花算法生成" json:"user_id"`
	WordType        int8           `gorm:"column:word_type;type:tinyint(1);not null;DEFAULT:0;comment:单词类型，默认为0，1为四级词汇，2为六级词汇，3为英专四级，4为英专八级，5为考研词汇，6为GRE词汇，7为托福词汇，8为雅思词汇" json:"word_type"`
	JustLearned     int8           `gorm:"column:just_learned;type:tinyint(1);not null;DEFAULT:0;comment:是否是刚学过，默认为0" json:"just_learned"`
	IsNeedLearned   int8           `gorm:"column:is_need_learned;type:tinyint(1);not null;DEFAULT:0;comment:是否需要学习，默认为0" json:"is_need_learned"`
	IsLearned       int8           `gorm:"column:is_learned;type:tinyint(1);not null;DEFAULT:0;comment:是否学习过，默认为0" json:"is_learned"`
	ExamNum         uint64         `gorm:"column:exam_num;type:bigint(20) unsigned;not null;DEFAULT:0;comment:总计检验次数" json:"exam_num"`
	ExamRightNum    uint64         `gorm:"column:exam_right_num;type:bigint(20) unsigned;not null;DEFAULT:0;comment:总计检验答对次数" json:"exam_right_num"`
	LastMasterTime  int64          `gorm:"column:last_master_time;type:bigint;not null;DEFAULT:0;comment:上次已掌握时间（时间戳）" json:"last_master_time"`
	LastReviewTime  int64          `gorm:"column:last_review_time;type:bigint;not null;DEFAULT:0;comment:上次复习的时间（时间戳）" json:"last_review_time"`
	MasterDegree    int            `gorm:"column:master_degree;type:int;not null;DEFAULT:0;comment:掌握程度（总计10分）" json:"master_degree"`
	DeepMasterTimes int            `gorm:"column:deep_master_times;type:int;not null;DEFAULT:0;comment:深度掌握次数" json:"deep_master_times"`
	NeedLearnDate   int64          `gorm:"column:need_learn_date;type:bigint;comment:需要学习的时间（以天为单位）" json:"need_learn_date"`
	NeedReviewDate  int64          `gorm:"column:need_review_date;type:bigint;comment:需要复习的时间（以天为单位）" json:"need_review_date"`
}
