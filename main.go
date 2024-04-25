package main

import (
	"bufio"
	"fmt"
	"my-palworld/routers"
	"os"
	"os/exec"
	"strings"
)

// 判断文件是否存在
func fileNotExist(filename string) bool {
	// 使用os.Stat()函数获取文件信息
	_, err := os.Stat(filename)

	return os.IsNotExist(err)
}

// 暂停程序
func pause() {
	fmt.Print("按任意键继续...")
	// 调用 cmd /c pause 命令来暂停程序执行
	cmd := exec.Command("cmd", "/c", "pause")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}
func checkServerPath() bool {
	// 暂时写死
	palServerExe := "PalServer.exe"
	// 如果文件不存在，提示用户
	if fileNotExist(palServerExe) {
		fmt.Printf("文件 %s 不存在\n", palServerExe)

		pause()
		return false
	} else {
		fmt.Printf("文件 %s 存在\n", palServerExe)
	}
	return true
}
func startWebUI() {
	//1.创建路由
	r := routers.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}

func readGameConfigFile() {

	// 配置位置
	// 先查找游戏自定义配置，如果有则读取内容，读取到匹配行后修改匹配到的数据，如果没有则使用默认配置
	// 读取当前目录下的文件 DefaultPalWorldSettings.ini
	file, err := os.Open("example.ini")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个新的 bufio 读取器
	scanner := bufio.NewScanner(file)

	// 用于存储修改后的文件内容
	var modifiedLines []string

	// 循环读取文件的每一行并进行处理
	for scanner.Scan() {
		line := scanner.Text()
		// 根据匹配条件修改行内容
		if strings.HasPrefix(line, "OptionSettings=(") {
			fmt.Println("修改前的内容: ")
			fmt.Println(line)
			// 去除开头和结尾

		}
		// 将修改后的行添加到切片中
		modifiedLines = append(modifiedLines, line)
	}

	// 检查扫描过程是否出错
	if err := scanner.Err(); err != nil {
		fmt.Println("扫描文件出错:", err)
		return
	}

	// 打开文件以便写入（截断原文件）
	outputFile, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer outputFile.Close()

	// 创建一个新的 bufio 写入器
	writer := bufio.NewWriter(outputFile)

	// 将修改后的内容写入文件
	for _, line := range modifiedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("写入文件出错:", err)
			return
		}
	}

	// 刷新缓冲区以确保所有数据都写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区出错:", err)
		return
	}

	fmt.Println("文件修改成功")
}
func main() {

	readGameConfigFile()
	// 获取系统信息
	//go system.GetSystemInfo()
	// 根据配置初始化之类。。。暂时跳过
	// 找到服务器目录（从配置中读取/手动输入）
	// 获取游戏状态
	//game.GetGameStatus()
	//
	//for {
	//	fmt.Println("请输入数字选择操作：")
	//	fmt.Println("1. 启动游戏")
	//	fmt.Println("2. 关闭游戏")
	//	fmt.Println("3. 重启游戏")
	//	fmt.Println("4. 更新游戏")
	//	fmt.Println("5. 退出")
	//
	//	var choice int
	//	_, err := fmt.Scanln(&choice)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//
	//	switch choice {
	//	case 1:
	//		err = game.StartGame()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	case 2:
	//		err = game.StopGame()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	case 3:
	//		err = game.RestartGame()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	case 4:
	//		err = game.UpdateGame()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	case 5:
	//		fmt.Println("程序退出")
	//		return
	//	default:
	//		fmt.Println("无效的选择，请重新输入。")
	//	}
	//}

	// todo 读取配置，启动停止重启服务器操作
	// 读取配置 后面做
	// 启动web页面
	//startWebUI()
	// todo 制作web页面，加入gin框架

	// 获取到当前路径

	// 控制客户端
	// 通过cmd控制

	// 启动

	// 关闭
	// 重启
	// 获取状态  根据系统不同使用tasklist获取（linux使用ps）帕鲁进程pid。。。
	//

	// 配置修改

	// web 相关

}
