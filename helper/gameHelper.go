package helper

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"my-palworld/config"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 根据进程名获取id
func GetProcessIdByProcessName(processName string) int {
	// 获取当前运行的进程列表
	fmt.Println("search process name: ", processName)
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	processId := 0
	for _, p := range processes {
		if strings.Contains(strings.ToLower(p.Executable()), processName) {
			processId = p.Pid()
			fmt.Println(processId)
			break
		}
	}

	return processId

}

// 根据进程名停止进程
func KillProcessByProcessName(processName string) error {
	// 使用 taskkill 命令停止进程（在 Windows 下）
	cmd := exec.Command("taskkill", "/f", "/im", processName)

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

// 根据app id 使用steamcmd更新游戏
func UpdateGameBySteamCmd(appId string) error {
	cmd := exec.Command(config.STEAM_CMD_PATH+"\\steamcmd.exe", "+login", "anonymous", "+app_update", appId, "validate", "+quit")
	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func CheckGameFile() bool {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// 遍历当前目录中的文件
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 获取文件名
			filename := filepath.Base(path)
			// 检查文件名是否包含 "aaa.exe"
			if strings.Contains(filename, config.GAME_NAME) {
				// 退出遍历
				fmt.Printf("Found file named %s at: %s\n", config.GAME_NAME, path)
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}
