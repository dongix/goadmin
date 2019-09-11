package web

import (
	"net/http"

	"dao"

	"service"

	"github.com/gin-gonic/gin"
)

// GetUserByID 根据用户id，查询用户
func GetUserByID(c *gin.Context) {
	var userQuery UserQuery
	if error := c.ShouldBindUri(&userQuery); error != nil {
		c.JSON(400, gin.H{"msg": error.Error()})
		return
	}
	user, error := dao.GetUserByID(userQuery.ID)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": error.Error(),
		})
	}
	c.JSON(http.StatusOK, user)
}

// GenSMSCode 生成验证码
func GenSMSCode(c *gin.Context) {
	var smscode SmsCode
	if error := c.ShouldBindUri(&smscode); error != nil {
		c.JSON(400, gin.H{"msg": error.Error()})
		return
	}
	error := service.GenSMSCode(smscode.Mobile)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Gen SMS code success!",
	})
}
