package model

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID                uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:用户信息表主键" json:"id"`
	CreatedAt         time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	UserId            uint64         `gorm:"column:user_id;index;type:bigint(20) unsigned;comment:用户id，雪花算法生成" json:"user_id"`
	LeXueAppId        string         `gorm:"column:le_xue_app_id;type:varchar(128);comment:交友id，也是唯一标识，初始为手机号" json:"le_xue_app_id"`
	Gender            int8           `gorm:"column:gender;type:tinyint(1);not null;DEFAULT:0;comment:用户性别，1为男性，2为女性，0为未知，默认为0" json:"gender"`
	School            string         `gorm:"column:school;type:varchar(255);comment:学校" json:"school"`
	Birthday          time.Time      `gorm:"column:birthday;type:datetime(3);comment:生日" json:"birthday"`
	Area              string         `gorm:"column:area;type:varchar(255);comment:地区" json:"area"`
	NickName          string         `gorm:"column:nickname;type:varchar(255);comment:昵称，注册后为随机" json:"nickname"`
	HeadSculpture     string         `gorm:"column:head_sculpture;type:varchar(255);comment:头像云存储地址" json:"head_sculpture"`
	Integral          uint64         `gorm:"column:integral;type:bigint(20) unsigned;not null;DEFAULT:0;comment:积分，默认为0，无符号" json:"integral"`
	WordNeedReciteNum int            `gorm:"column:word_need_recite_num;type:int;not null;DEFAULT:0;comment:每日需要背单词的数量" json:"word_need_recite_num"`
	EngLevel          int8           `gorm:"column:eng_level;type:tinyint(1);not null;DEFAULT:0;comment:词书等级，默认为0。1为四级，2为六级等" json:"eng_level"`
	LastStartTime     time.Time      `gorm:"column:last_start_time;type:datetime(3);comment:上次学习的时间" json:"last_start_time"`
	Role              int8           `gorm:"column:role;type:tinyint(1);not null;DEFAULT:0;comment:用户权限，0为正常，1为VIP，默认为0" json:"role"`
	InvitationCode    string         `gorm:"column:invitation_code;type:varchar(32);comment:邀请码，独一无二的邀请码" json:"invitation_code"`
	Signature         string         `gorm:"column:signature;type:varchar(255);comment:个性签名" json:"signature"`
	DeleteIsOk        int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:用户信息是否删除，默认为0，1是删除" json:"delete_isok"`
}
