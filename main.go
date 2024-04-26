package main

import (
	"fmt"
	"my-palworld/helper"
	"my-palworld/routers"
	"os"
	"os/exec"
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
	err := cmd.Run()
	if err != nil {
		return
	}
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

func initConfig() error {
	// 初始化配置文件
	err := helper.ReadConfigFile()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	err = helper.ReadGameConfigFile()
	if err != nil {
		return err
	}
	return nil
}
func main() {

	err := initConfig()
	if err != nil {
		fmt.Printf("init config err: %s\n", err.Error())
		return
	}

	// 读取配置

	//
	//// 写入配置
	//err := helper.WriteGameConfigFile()
	//if err != nil {
	//	fmt.Printf(err.Error())
	//	return
	//}
	// 获取系统信息
	//go system.GetSystemInfo()
	// 根据配置初始化之类。。。暂时跳过
	// 找到服务器目录（从配置中读取/手动输入）

	// todo 读取配置，启动停止重启服务器操作
	// 读取配置 后面做
	// 启动web页面
	startWebUI()
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
