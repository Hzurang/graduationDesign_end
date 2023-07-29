package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:用户表主键" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	UserId        uint64         `gorm:"column:user_id;index;type:bigint(20) unsigned;comment:用户id，雪花算法生成，无符号" json:"user_id"`
	Email         string         `gorm:"column:email;type:varchar(64);comment:邮箱" json:"email" binding:"email"`
	Mobile        string         `gorm:"column:mobile;type:varchar(64);comment:手机号" json:"mobile"`
	Password      string         `gorm:"column:password;type:varchar(255);comment:密码" json:"password"`
	Status        int8           `gorm:"column:status;type:tinyint(1);not null;DEFAULT:0;comment:用户状态，0为正常，1为禁用，默认为0" json:"status"`
	LastLoginTime time.Time      `gorm:"column:last_login_time;type:datetime(3);comment:上次登录时间" json:"last_login_time"`
}

type LoginUser struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	IP       string `json:"ip"`
	AcToken  string `json:"ac_token"`
	ReToken  string `json:"re_token"`
	EngLevel string `json:"eng_level"`
}
