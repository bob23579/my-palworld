package game

import (
	"fmt"
	"my-palworld/config"
	"my-palworld/helper"
	"strings"
)

// 获取游戏状态
func GetGameStatus() {

	// 查找进程中是否在运行cmd版进程
	pid := helper.GetProcessIdByProcessName(strings.ToLower(config.GAME_CMD_NAME))
	if pid != 0 {
		config.GameStatus = config.GameStart
		fmt.Println("find game cmd process success")
		return
	}
	// 如果不在，查找steam cmd进程，如果在，返回更新
	pid = helper.GetProcessIdByProcessName("steamcmd.exe")
	if pid != 0 {
		config.GameStatus = config.GameUpdating
		return
	}
	//如果不在，查找目录中是否有游戏程序，如果有返回停止
	if helper.CheckGameFile() {
		config.GameStatus = config.GameStop
		return
	}
	// 如果没有，返回没找到程序
	// 重启状态根据命令手动设置，或者后期通过hook判断是否正在重启

	config.GameStatus = config.GameNotFound
}
