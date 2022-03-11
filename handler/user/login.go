package user

import (
	"TreePlanting/handler"
	"TreePlanting/model"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login 登录
// @Summary  登录
// @Tags user
// @Description 学号密码登录
// @Accept application/json
// @Produce application/json
// @Param object body model.User true "登录的用户信息"
// @Success 200 {object} handler.Response "{"msg":"将student_id作为token保留"}"
// @Failure 401 {object} string "{"error_code":"10001", "message":"Password or account wrong."} 身份认证失败 重新登录"
// @Failure 400 {object} string "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} string "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user [post]
func Login(c *gin.Context) {
	var p model.User
	if err := c.BindJSON(&p); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if len(p.StudentID) != 10 || p.StudentID[0] != '2' {
		c.JSON(400, gin.H{"message": "student_id Not Satisfiable."})
		return
	}
	if p.Password == "" {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	user := model.UserModel{
		StudentID: p.StudentID,
	}
	if res := model.DB.Where("student_id = ?", p.StudentID).First(&p); res.Error != nil {

		model.DB.Table("users").Create(&user)
	} else {
		// 验证一站式
		_, err := model.GetUserInfoFormOne(p.StudentID, p.Password)
		if err != nil {
			fmt.Println(err)
			c.JSON(401, "Password or account wrong.")
			return
		}
	}

	claims := &model.Jwt{StudentID: p.StudentID}

	claims.ExpiresAt = time.Now().Add(200 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	var Secret = "vinegar" // 加醋

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
	}

	handler.SendResponse(c, "将student_id作为token保留", signedToken)
}
