package model

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 全局变量
var DB *gorm.DB

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Errorf("Open database failed, %s\n", err.Error())
		panic(err)
	}

	return db
}

func InitDB()  {
	DB = openDB(viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.addr"), viper.GetString("db.name"))
}

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

func GetContent(user UserModel) (string, error) {
	err := DB.Table("users").Where("student_id = ?", user.StudentID).First(&user).Error
	return user.StudentID, err
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
