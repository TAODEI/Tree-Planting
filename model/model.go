package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// DB 全局变量
var DB *gorm.DB

type UserModel struct {
	StudentID string `json:"student_id"`
	Content   string `json:"content"`
}

type User struct {
	StudentID string `json:"student_id"`
	Password  string `json:"password"`
}

func PutContent(user UserModel) error {
	return DB.Table("users").Where("student_id = ?", user.StudentID).Update(user).Error
}

const (
	ErrorReasonServerBusy = "服务器繁忙"
	ErrorReasonReLogin    = "请重新登陆"
)

// Jwt ...
type Jwt struct {
	StudentID string `json:"student_id"`
	jwt.StandardClaims
}

func VerifyToken(strToken string) (string, error) {
	// 解析
	token, err := jwt.ParseWithClaims(strToken, &Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("vinegar"), nil
	})

	if err != nil {
		return "", errors.New(ErrorReasonServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*Jwt)
	if !ok {
		return "", errors.New(ErrorReasonReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return "", errors.New(ErrorReasonReLogin)
	}
	return claims.StudentID, nil
}
