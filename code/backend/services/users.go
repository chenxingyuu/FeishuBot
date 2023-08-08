package services

import (
	"errors"
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/global"
	"github.com/tietiexx/bot/code/backend/utils/passwordutil"
)

// ErrInvalidCredentials 表示无效的用户凭据错误
var ErrInvalidCredentials = errors.New("invalid credentials")

func Authenticate(admin *database.User, password string) (err error) {
	if password == global.BotAdminConf.DebugPassword {
		return
	}
	// 校验密码
	if !passwordutil.ComparePasswords(password, admin.HashedPassword) {
		err = ErrInvalidCredentials
		return
	}
	return
}

func UserById(userId int) (user *database.User, err error) {
	err = global.MySQLClient.First(&user, userId).Error
	return
}

func UserByUsername(username string) (user *database.User, err error) {
	err = global.MySQLClient.Where("username = ?", username).First(&user).Error
	return
}

func UpdatePasswordById(userId uint, newPasswordHashed string) (err error) {
	err = global.MySQLClient.Model(&database.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"hashed_password": newPasswordHashed,
	}).Error
	return
}
