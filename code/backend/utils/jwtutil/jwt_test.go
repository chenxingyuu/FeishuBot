package jwtutil

import (
	"testing"
)

func TestGenerateAccessToken(t *testing.T) {
	userID := 1

	// 生成 JWT
	token, err := GenerateAccessToken(userID)
	if err != nil {
		t.Errorf("生成 JWT 失败: %v", err)
	}

	// 验证并解析 JWT
	claims, err := VerifyToken(token)
	if err != nil {
		t.Errorf("验证 JWT 失败: %v", err)
	}

	// 检查用户ID是否正确
	if claims.UserID != userID {
		t.Errorf("用户ID不匹配，期望 %d，实际 %d", userID, claims.UserID)
	}

	// 检查TokenType是否正确
	if claims.TokenType != Access {
		t.Errorf("TokenType不匹配，期望 %d，实际 %d", Access, claims.TokenType)
	}

}

func TestGenerateRefreshToken(t *testing.T) {
	userID := 1

	// 生成 JWT
	token, err := GenerateRefreshToken(userID)
	if err != nil {
		t.Errorf("生成 JWT 失败: %v", err)
	}

	// 验证并解析 JWT
	claims, err := VerifyToken(token)
	if err != nil {
		t.Errorf("验证 JWT 失败: %v", err)
	}

	// 检查用户ID是否正确
	if claims.UserID != userID {
		t.Errorf("用户ID不匹配，期望 %d，实际 %d", userID, claims.UserID)
	}

	// 检查TokenType是否正确
	if claims.TokenType != Refresh {
		t.Errorf("TokenType不匹配，期望 %d，实际 %d", Refresh, claims.TokenType)
	}

}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODY4ODM2MzcsImlhdCI6MTY4Njg4MzYxM30.Um6OmpknnJLpnny8jAdgZiTFWcx9OxxaRO5A-Qn-_6o"

	// 验证并解析 JWT
	_, err := VerifyToken(token)
	if err == nil {
		t.Error("未检测到 JWT 过期错误")
	}

	// 检查过期错误信息
	expectedError := "token已过期"
	if err.Error() != expectedError {
		t.Errorf("错误信息不匹配，期望 %s，实际 %s", expectedError, err.Error())
	}
}
