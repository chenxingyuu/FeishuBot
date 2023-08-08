package database

import (
	"github.com/tietiexx/bot/code/backend/constant"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	BaseModel
	Username       string `gorm:"unique;comment:用户名称" json:"username"`
	HashedPassword string `gorm:"comment:用户密码" json:"hashed_password"`
	Status         int8   `gorm:"default:1;comment:用户状态 0-冻结 1-正常" json:"status"`
}

type LarkApp struct {
	BaseModel
	UUID              string  `gorm:"comment:bot uuid;size:32" json:"uuid"`
	Name              string  `gorm:"comment:bot name;size:32" json:"name"`
	AppId             string  `gorm:"size:32" json:"app_id"`
	AppSecret         string  `gorm:"size:32" json:"app_secret"`
	EncryptKey        string  `gorm:"size:32" json:"encrypt_key"`
	VerificationToken string  `gorm:"size:32" json:"verification_token"`
	LarkBot           LarkBot `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type LarkBot struct {
	BaseModel
	UUID      string               `gorm:"comment:bot uuid;size:32" json:"uuid"`
	Name      string               `gorm:"comment:bot name;size:32" json:"name"`
	BotType   constant.LakeBotType `gorm:"comment:机器人类型 1-自定义机器人 2-应用机器人" json:"bot_type"`
	LarkAppID uint
}
