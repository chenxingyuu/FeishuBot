package passwordutil

import (
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := "password123"

	// 加密密码
	hashedPassword, err := EncryptPassword(password)
	if err != nil {
		t.Errorf("加密密码失败: %v", err)
	}

	// 检查加密后的密码是否与原始密码不同
	if hashedPassword == password {
		t.Errorf("加密后的密码与原始密码相同")
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password123"
	hashedPassword, _ := EncryptPassword(password)

	// 验证正确的密码
	match := ComparePasswords(password, hashedPassword)
	if !match {
		t.Error("正确的密码验证失败")
	}

	// 验证不正确的密码
	invalidPassword := "wrong_password"
	match = ComparePasswords(invalidPassword, hashedPassword)
	if match {
		t.Error("不正确的密码验证通过")
	}
}
