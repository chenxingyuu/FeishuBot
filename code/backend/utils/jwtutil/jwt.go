package jwtutil

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secretKey = []byte("xx") // 自定义密钥，用于签名和验证
)

const (
	RefreshExpiration = time.Hour * 24 * 30
	AccessExpiration  = time.Hour * 24 * 2
)

type TokenType int

const (
	Refresh TokenType = iota
	Access
)

type CustomClaims struct {
	UserID    int       `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	jwt.StandardClaims
}

func GenerateAccessToken(userID int) (string, error) {
	claims := CustomClaims{
		UserID:    userID,
		TokenType: Access,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	return GenerateToken(claims)
}

func GenerateRefreshToken(userID int) (string, error) {
	claims := CustomClaims{
		UserID:    userID,
		TokenType: Refresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	return GenerateToken(claims)
}

// GenerateToken 生成 JWT
func GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken 验证并解析 JWT
func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的Token")
}

func VerifyAccessToken(tokenString string) (*CustomClaims, error) {
	claims, err := VerifyToken(tokenString)
	if err == nil && claims.TokenType == Access {
		return claims, nil
	}
	return nil, err
}

func VerifyRefreshToken(tokenString string) (*CustomClaims, error) {
	claims, err := VerifyToken(tokenString)
	if err == nil && claims.TokenType == Refresh {
		return claims, nil
	}
	return nil, err
}
