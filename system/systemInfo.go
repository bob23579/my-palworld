package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"my-palworld/config"
	"time"
)
import "github.com/shirou/gopsutil/mem"

func GetSystemInfo() {
	// 将状态信息放入数组中
	var systemInfoList []config.SystemInfo
	for range time.Tick(3 * time.Second) {
		var systemInfo config.SystemInfo
		// 获取当前时间（精确到秒）
		// 获取当前时间
		currentTime := time.Now()
		// 格式化为精确到秒的字符串
		timeString := currentTime.Format("2006-01-02 15:04:05")

		systemInfo.Time = timeString
		// 输出格式化后的时间字符串
		fmt.Println("Current time :", timeString)
		// 获取系统内存使用情况
		memInfo, _ := mem.VirtualMemory()

		// 打印系统内存占比
		fmt.Printf("Memory Usage: %.2f%%\n", memInfo.UsedPercent)
		formatted := fmt.Sprintf("%.2f", memInfo.UsedPercent)
		systemInfo.MemoryList = formatted
		// 获取 CPU 使用情况
		cpuUsage, _ := cpu.Percent(0, false)
		formatted = fmt.Sprintf("%.2f", cpuUsage[0])
		systemInfo.CpuList = formatted
		fmt.Printf("cpu Usage: %.2f%%\n", cpuUsage[0])
		// 获取硬盘使用情况
		diskUsage, _ := disk.Usage("/")
		fmt.Printf("disk Usage: %.2f%%\n", diskUsage.UsedPercent)
		formatted = fmt.Sprintf("%.2f", diskUsage.UsedPercent)
		systemInfo.DistList = formatted

		systemInfoList = append(systemInfoList, systemInfo)

	}
}
