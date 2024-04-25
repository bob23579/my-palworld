package game

import (
	"errors"
	"fmt"
	"my-palworld/config"
	"my-palworld/helper"
	"os/exec"
)

// StartGame 启动游戏
func StartGame() error {
	// 检查是否已经启动了，如果已经启动了，返回error
	if config.GameStatus == config.GameStart {
		return errors.New("the game is already running")
	}
	// 启动游戏
	// "\Pal\Binaries\Win64\PalServer-Win64-Shipping-Cmd.exe"
	// 当前目录下的pal server文件
	cmd := exec.Command(".\\" + config.GAME_NAME)

	// 启动程序
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// 设置状态

	config.GameStatus = config.GameStart
	return nil
}

// 关闭游戏
func StopGame() error {
	// 如果游戏已经关闭，则直接返回
	if config.GameStatus == config.GameStop {
		return nil
	}

	err := helper.KillProcessByProcessName(config.GAME_CMD_NAME)
	if err != nil {
		return err
	}
	// 要停止的进程名
	//err = helper.KillProcessByProcessName(config.GAME_NAME)
	//if err != nil {
	//	return err
	//}

	fmt.Printf("Process '%s' stopped\n", config.GAME_NAME)
	config.GameStatus = config.GameStop
	return nil
}

// 重启游戏
func RestartGame() error {
	// 判断游戏是否已经启动，如果启动了 先关闭再启动
	if config.GameStatus == config.GameStart {
		err := StopGame()
		if err != nil {
			return err
		}
	}
	// 否则直接启动
	err := StartGame()
	if err != nil {
		return err
	}
	return err
}

// 更新游戏
func UpdateGame() error {
	// 判断游戏是否启动
	if config.GameStatus == config.GameStart {
		err := StopGame()
		if err != nil {
			return err
		}
	}
	// 先关闭游戏再更新
	err := helper.UpdateGameBySteamCmd(config.PalSteamCmdID)
	if err != nil {
		return err
	}

	// 更新命令 appid var PalSteamCmdID = 2394010
	// cmd := exec.Command(steamcmdPath, "+login", "anonymous", "+app_update", appID, "validate", "+quit")
	return nil
}

// 保存存档。。。。回档。。。。清理存档等
