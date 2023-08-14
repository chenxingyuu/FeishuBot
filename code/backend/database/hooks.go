package database

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("创建用户")
	return err
}

func (a *LarkApp) BeforeCreate(tx *gorm.DB) (err error) {
	a.UUID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return err
}

func (b *LarkBot) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return err
}
