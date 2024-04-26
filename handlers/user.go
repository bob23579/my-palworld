package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-palworld/config"
	"my-palworld/helper"
	"net/http"
)

type LoginRequestData struct {
	// 定义前端发送的登录请求数据结构
	Username string `json:"username"`
	Password string `json:"password"`
	// 其他字段...
}

type PasswordChange struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func LoginHandler(c *gin.Context) {
	// 判断输入参数
	var requestData LoginRequestData
	err := c.BindJSON(&requestData)
	if err != nil {
		fmt.Println(requestData)
		return
	}
	name := requestData.Username
	password := requestData.Password
	fmt.Println(name)
	fmt.Println(password)

	// 用户名和密码暂时写死，后期修改
	if name == config.EditConfig.Username && password == config.EditConfig.Password {
		// 生成token返回
		token, err := generateToken(name)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"token": token, "message": "success"}})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"message": "username or password error"}})
	}
}

func GetUserInfo(c *gin.Context) {
	name, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"username": name, "roles": []string{"admin"}}})
}

func ModifyPassword(c *gin.Context) {
	// 判断旧密码是否正确,不正确返回,正确使用新密码
	// 打印传入的数据
	c.Param("oldPassword")
	var data PasswordChange
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101, "message": "modify password error : " + err.Error()})
		return
	}

	if data.OldPassword != config.EditConfig.Password {
		c.JSON(http.StatusOK, gin.H{"code": 101, "message": "old password error"})
		return
	}
	// 修改原密码
	config.EditConfig.Password = data.NewPassword
	// 写入文件
	err = helper.WriteConfigFile()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101, "message": "error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success"})
}
