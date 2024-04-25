package helper

import (
	"bufio"
	"fmt"
	"my-palworld/config"
	"os"
	"strconv"
	"strings"
)

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Error parsing float: %v\n", err)
		return 0
	}
	return f
}
func parseBool(s string) bool {
	f, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Printf("Error parsing float: %v\n", err)
		return false
	}
	return f
}
func parseInt(s string) int {
	f, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error parsing float: %v\n", err)
		return 0
	}
	return f
}

// 读取游戏配置文件 对照进结构体，不修改文件
func ReadGameConfigFile() error {

	// 配置位置
	// 先查找游戏自定义配置，如果有则读取内容，读取到匹配行后修改匹配到的数据，如果没有则使用默认配置
	// 读取当前目录下的文件 DefaultPalWorldSettings.ini
	file, err := os.Open("example.ini")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// 创建一个新的 bufio 读取器
	scanner := bufio.NewScanner(file)

	// 用于存储修改后的文件内容
	var modifiedLines []string

	// 循环读取文件的每一行并进行处理
	for scanner.Scan() {
		line := scanner.Text()
		// 根据匹配条件修改行内容
		if strings.HasPrefix(line, "OptionSettings=(") {
			line = line[16 : len(line)-1]
			// 根据逗号切割,对照结构体
			modifiedLines = strings.Split(line, ",")
			break
		}

	}
	// 映射到结构体中
	for _, pair := range modifiedLines {
		keyValue := strings.Split(pair, "=")
		key := keyValue[0]
		value := keyValue[1]
		// 去除字符串中的双引号
		value = strings.Replace(value, "\"", "", -1)

		switch key {
		case "Difficulty":
			config.EditGameConfig.Difficulty = value
		case "DayTimeSpeedRate":
			config.EditGameConfig.DayTimeSpeedRate = parseFloat(value)
		case "NightTimeSpeedRate":
			config.EditGameConfig.NightTimeSpeedRate = parseFloat(value)
		case "ExpRate":
			config.EditGameConfig.ExpRate = parseFloat(value)
		case "PalCaptureRate":
			config.EditGameConfig.PalCaptureRate = parseFloat(value)
		case "PalSpawnNumRate":
			config.EditGameConfig.PalSpawnNumRate = parseFloat(value)
		case "PalDamageRateAttack":
			config.EditGameConfig.PalDamageRateAttack = parseFloat(value)
		case "PalDamageRateDefense":
			config.EditGameConfig.PalDamageRateDefense = parseFloat(value)
		case "PlayerDamageRateAttack":
			config.EditGameConfig.PlayerDamageRateAttack = parseFloat(value)
		case "PlayerDamageRateDefense":
			config.EditGameConfig.PlayerDamageRateDefense = parseFloat(value)
		case "PlayerStomachDecreaceRate":
			config.EditGameConfig.PlayerStomachDecreaceRate = parseFloat(value)
		case "PlayerStaminaDecreaceRate":
			config.EditGameConfig.PlayerStaminaDecreaceRate = parseFloat(value)
		case "PlayerAutoHPRegeneRate":
			config.EditGameConfig.PlayerAutoHPRegeneRate = parseFloat(value)
		case "PlayerAutoHpRegeneRateInSleep":
			config.EditGameConfig.PlayerAutoHpRegeneRateInSleep = parseFloat(value)
		case "PalStomachDecreaceRate":
			config.EditGameConfig.PalStomachDecreaceRate = parseFloat(value)
		case "PalStaminaDecreaceRate":
			config.EditGameConfig.PalStaminaDecreaceRate = parseFloat(value)
		case "PalAutoHPRegeneRate":
			config.EditGameConfig.PalAutoHPRegeneRate = parseFloat(value)
		case "PalAutoHpRegeneRateInSleep":
			config.EditGameConfig.PalAutoHpRegeneRateInSleep = parseFloat(value)
		case "BuildObjectDamageRate":
			config.EditGameConfig.BuildObjectDamageRate = parseFloat(value)
		case "BuildObjectDeteriorationDamageRate":
			config.EditGameConfig.BuildObjectDeteriorationDamageRate = parseFloat(value)
		case "CollectionDropRate":
			config.EditGameConfig.CollectionDropRate = parseFloat(value)
		case "CollectionObjectHpRate":
			config.EditGameConfig.CollectionObjectHpRate = parseFloat(value)
		case "CollectionObjectRespawnSpeedRate":
			config.EditGameConfig.CollectionObjectRespawnSpeedRate = parseFloat(value)
		case "EnemyDropItemRate":
			config.EditGameConfig.EnemyDropItemRate = parseFloat(value)
		case "DeathPenalty":
			config.EditGameConfig.DeathPenalty = value
		case "bEnablePlayerToPlayerDamage":
			config.EditGameConfig.BEnablePlayerToPlayerDamage = parseBool(value)
		case "bEnableFriendlyFire":
			config.EditGameConfig.BEnableFriendlyFire = parseBool(value)
		case "bEnableInvaderEnemy":
			config.EditGameConfig.BEnableInvaderEnemy = parseBool(value)
		case "bActiveUNKO":
			config.EditGameConfig.BActiveUNKO = parseBool(value)
		case "bEnableAimAssistPad":
			config.EditGameConfig.BEnableAimAssistPad = parseBool(value)
		case "bEnableAimAssistKeyboard":
			config.EditGameConfig.BEnableAimAssistKeyboard = parseBool(value)
		case "DropItemMaxNum":
			config.EditGameConfig.DropItemMaxNum = parseInt(value)
		case "DropItemMaxNum_UNKO":
			config.EditGameConfig.DropItemMaxNum_UNKO = parseInt(value)
		case "BaseCampMaxNum":
			config.EditGameConfig.BaseCampMaxNum = parseInt(value)
		case "BaseCampWorkerMaxNum":
			config.EditGameConfig.BaseCampWorkerMaxNum = parseInt(value)
		case "DropItemAliveMaxHours":
			config.EditGameConfig.DropItemAliveMaxHours = parseFloat(value)
		case "bAutoResetGuildNoOnlinePlayers":
			config.EditGameConfig.BAutoResetGuildNoOnlinePlayers = parseBool(value)
		case "AutoResetGuildTimeNoOnlinePlayers":
			config.EditGameConfig.AutoResetGuildTimeNoOnlinePlayers = parseFloat(value)
		case "GuildPlayerMaxNum":
			config.EditGameConfig.GuildPlayerMaxNum = parseInt(value)
		case "PalEggDefaultHatchingTime":
			config.EditGameConfig.PalEggDefaultHatchingTime = parseFloat(value)
		case "WorkSpeedRate":
			config.EditGameConfig.WorkSpeedRate = parseFloat(value)
		case "bIsMultiplay":
			config.EditGameConfig.BIsMultiplay = parseBool(value)
		case "bIsPvP":
			config.EditGameConfig.BIsPvP = parseBool(value)
		case "bCanPickupOtherGuildDeathPenaltyDrop":
			config.EditGameConfig.BCanPickupOtherGuildDeathPenaltyDrop = parseBool(value)
		case "bEnableNonLoginPenalty":
			config.EditGameConfig.BEnableNonLoginPenalty = parseBool(value)
		case "bEnableFastTravel":
			config.EditGameConfig.BEnableFastTravel = parseBool(value)
		case "bIsStartLocationSelectByMap":
			config.EditGameConfig.BIsStartLocationSelectByMap = parseBool(value)
		case "bExistPlayerAfterLogout":
			config.EditGameConfig.BExistPlayerAfterLogout = parseBool(value)
		case "bEnableDefenseOtherGuildPlayer":
			config.EditGameConfig.BEnableDefenseOtherGuildPlayer = parseBool(value)
		case "CoopPlayerMaxNum":
			config.EditGameConfig.CoopPlayerMaxNum = parseInt(value)
		case "ServerPlayerMaxNum":
			config.EditGameConfig.ServerPlayerMaxNum = parseInt(value)
		case "ServerName":
			config.EditGameConfig.ServerName = value
		case "ServerDescription":
			config.EditGameConfig.ServerDescription = value
		case "AdminPassword":
			config.EditGameConfig.AdminPassword = value
		case "ServerPassword":
			config.EditGameConfig.ServerPassword = value
		case "PublicPort":
			config.EditGameConfig.PublicPort = parseInt(value)
		case "PublicIP":
			config.EditGameConfig.PublicIP = value
		case "RCONEnabled":
			config.EditGameConfig.RCONEnabled = parseBool(value)
		case "RCONPort":
			config.EditGameConfig.RCONPort = parseInt(value)
		case "Region":
			config.EditGameConfig.PublicIP = value
		case "bUseAuth":
			config.EditGameConfig.BUseAuth = parseBool(value)
		case "BanListURL":
			config.EditGameConfig.BanListURL = value
		case "RESTAPIEnabled":
			config.EditGameConfig.RESTAPIEnabled = parseBool(value)
		case "RESTAPIPort":
			config.EditGameConfig.RESTAPIPort = parseInt(value)
		case "bShowPlayerList":
			config.EditGameConfig.BShowPlayerList = parseBool(value)
		case "AllowConnectPlatform":
			config.EditGameConfig.AllowConnectPlatform = value
		case "bIsUseBackupSaveData":
			config.EditGameConfig.BIsUseBackupSaveData = parseBool(value)
		case "LogFormatType":
			config.EditGameConfig.LogFormatType = value
		}
	}

	return nil
}

// 写入游戏配置文件
func WriteGameConfigFile() error {
	// 首先把前面部分写入
	// 要写入的文件名
	filename := "example-new.txt"

	// 打开文件，如果文件不存在则创建，如果文件已存在则截断
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建文件时发生错误:", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// 创建写入器
	writer := bufio.NewWriter(file)

	// 写入文件的每一行
	lines := []string{
		"[/Script/Pal.PalGameWorldSettings]",
	}
	lines = append(lines, config.EditGameConfig.ToString())

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("写入文件时发生错误:", err)
			return err
		}
	}

	// 刷新缓冲区，确保所有数据写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区时发生错误:", err)
		return err
	}

	fmt.Println("文件写入成功！")
	return nil
}

// 读取后台配置文件

// 写入后台配置文件
