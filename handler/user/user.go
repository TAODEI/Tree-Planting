package user

import (
	"TreePlanting/handler"
	"TreePlanting/model"
	"github.com/gin-gonic/gin"
)

// PushContent
// @Summary  上传content
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body Content true "content"
// @Success 200 {object} handler.Response "{"msg":"修改成功"}"
// @Failure 401 {object} string "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} string "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} string "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /content [post]
func PushContent(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}
	var content Content
	if err := c.BindJSON(&content); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	user := model.UserModel{
		StudentID: id,
		Content:   content.Text,
	}

	if err := model.PutContent(user); err != nil {
		c.JSON(400, gin.H{"message": "更新失败"})
		return
	}
	handler.SendResponse(c, "修改成功", nil)
}

type Content struct {
	Text string `json:"text"`
}
