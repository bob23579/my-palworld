package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequestData struct {
	// 定义前端发送的登录请求数据结构
	Username string `json:"username"`
	Password string `json:"password"`
	// 其他字段...
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
	if name == "admin" && password == "12345678" {
		// 生成token返回
		token, err := generateToken(name)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"token": token}})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "success": false, "username": name, "password": password})
	}
}

func GetUserInfo(c *gin.Context) {
	name, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"username": name, "roles": []string{"admin"}}})
}

func ModifyPassword(c *gin.Context) {
	//

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"username": "hello", "roles": []string{"admin"}}})
}
