package handlers

import (
	"github.com/gin-gonic/gin"
	"my-palworld/config"
	"my-palworld/game"
	"net/http"
)

func GetGameStatusHandler(c *gin.Context) {
	// 获取当前状态返回

	game.GetGameStatus()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"status": config.GameStatus}})
	// 用户名和密码暂时写死，后期修改

}
func StartGameHandler(c *gin.Context) {

	err := game.StartGame()
	if err != nil {

		c.JSON(http.StatusOK, gin.H{"code": 101, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success"})

}
func StopGameHandler(c *gin.Context) {

	err := game.StopGame()
	if err != nil {

		c.JSON(http.StatusOK, gin.H{"code": 101, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success"})

}
func RestartGameHandler(c *gin.Context) {

	err := game.RestartGame()
	if err != nil {

		c.JSON(http.StatusOK, gin.H{"code": 101, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success"})

}
func UpdateGameHandler(c *gin.Context) {

	err := game.UpdateGame()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success"})

}
