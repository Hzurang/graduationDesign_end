package model

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID         uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;autoIncrement:true;comment:管理员表主键" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
	AdminId    uint64         `gorm:"column:admin_id;index;type:bigint(20) unsigned;comment:管理员id，雪花算法生成" json:"admin_id"`
	UserName   string         `gorm:"column:user_name;type:varchar(32);comment:用户名" json:"user_name"`
	Password   string         `gorm:"column:password;type:varchar(255);comment:密码" json:"password"`
	DeleteIsOk int8           `gorm:"column:delete_isok;type:tinyint(1);not null;DEFAULT:0;comment:管理员是否删除，默认为0，1是删除" json:"delete_isok"`
}

type AdminUser struct {
	AdminId  uint64 `json:"admin_id"`
	UserName string `json:"user_name"`
	AcToken  string `json:"ac_token"`
	ReToken  string `json:"re_token"`
}
