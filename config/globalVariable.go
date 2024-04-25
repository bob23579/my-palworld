package config

import "fmt"

type SystemInfo struct {
	MemoryList string
	CpuList    string
	DistList   string
	Time       string
}
type GameConfig struct {
	Difficulty                           string
	DayTimeSpeedRate                     float64
	NightTimeSpeedRate                   float64
	ExpRate                              float64
	PalCaptureRate                       float64
	PalSpawnNumRate                      float64
	PalDamageRateAttack                  float64
	PalDamageRateDefense                 float64
	PlayerDamageRateAttack               float64
	PlayerDamageRateDefense              float64
	PlayerStomachDecreaceRate            float64
	PlayerStaminaDecreaceRate            float64
	PlayerAutoHPRegeneRate               float64
	PlayerAutoHpRegeneRateInSleep        float64
	PalStomachDecreaceRate               float64
	PalStaminaDecreaceRate               float64
	PalAutoHPRegeneRate                  float64
	PalAutoHpRegeneRateInSleep           float64
	BuildObjectDamageRate                float64
	BuildObjectDeteriorationDamageRate   float64
	CollectionDropRate                   float64
	CollectionObjectHpRate               float64
	CollectionObjectRespawnSpeedRate     float64
	EnemyDropItemRate                    float64
	DeathPenalty                         string
	BEnablePlayerToPlayerDamage          bool
	BEnableFriendlyFire                  bool
	BEnableInvaderEnemy                  bool
	BActiveUNKO                          bool
	BEnableAimAssistPad                  bool
	BEnableAimAssistKeyboard             bool
	DropItemMaxNum                       int
	DropItemMaxNum_UNKO                  int
	BaseCampMaxNum                       int
	BaseCampWorkerMaxNum                 int
	DropItemAliveMaxHours                float64
	BAutoResetGuildNoOnlinePlayers       bool
	AutoResetGuildTimeNoOnlinePlayers    float64
	GuildPlayerMaxNum                    int
	PalEggDefaultHatchingTime            float64
	WorkSpeedRate                        float64
	BIsMultiplay                         bool
	BIsPvP                               bool
	BCanPickupOtherGuildDeathPenaltyDrop bool
	BEnableNonLoginPenalty               bool
	BEnableFastTravel                    bool
	BIsStartLocationSelectByMap          bool
	BExistPlayerAfterLogout              bool
	BEnableDefenseOtherGuildPlayer       bool
	CoopPlayerMaxNum                     int
	ServerPlayerMaxNum                   int
	ServerName                           string
	ServerDescription                    string
	AdminPassword                        string
	ServerPassword                       string
	PublicPort                           int
	PublicIP                             string
	RCONEnabled                          bool
	RCONPort                             int
	Region                               string
	BUseAuth                             bool
	BanListURL                           string
	RESTAPIEnabled                       bool
	RESTAPIPort                          int
	BShowPlayerList                      bool
	AllowConnectPlatform                 string
	BIsUseBackupSaveData                 bool
	LogFormatType                        string
}

func getBoolString(b bool) string {
	if b {
		return "True"
	} else {
		return "False"
	}
}
func (s GameConfig) ToString() string {
	// 使用字符串拼接将结构体字段连接为一个字符串
	str := fmt.Sprintf("OptionSettings=(Difficulty=%s,DayTimeSpeedRate=%f,NightTimeSpeedRate=%f,ExpRate=%f,PalCaptureRate=%f,PalSpawnNumRate=%f,PalDamageRateAttack=%f,PalDamageRateDefense=%f,PlayerDamageRateAttack=%f,PlayerDamageRateDefense=%f,PlayerStomachDecreaceRate=%f,PlayerStaminaDecreaceRate=%f,PlayerAutoHPRegeneRate=%f,PlayerAutoHpRegeneRateInSleep=%f,PalStomachDecreaceRate=%f,PalStaminaDecreaceRate=%f,PalAutoHPRegeneRate=%f,PalAutoHpRegeneRateInSleep=%f,BuildObjectDamageRate=%f,BuildObjectDeteriorationDamageRate=%f,CollectionDropRate=%f,CollectionObjectHpRate=%f,CollectionObjectRespawnSpeedRate=%f,EnemyDropItemRate=%f,DeathPenalty=%s,bEnablePlayerToPlayerDamage=%s,bEnableFriendlyFire=%s,bEnableInvaderEnemy=%s,bActiveUNKO=%s,bEnableAimAssistPad=%s,bEnableAimAssistKeyboard=%s,DropItemMaxNum=%d,DropItemMaxNum_UNKO=%d,BaseCampMaxNum=%d,BaseCampWorkerMaxNum=%d,DropItemAliveMaxHours=%f,bAutoResetGuildNoOnlinePlayers=%s,AutoResetGuildTimeNoOnlinePlayers=%f,GuildPlayerMaxNum=%d,PalEggDefaultHatchingTime=%f,WorkSpeedRate=%f,bIsMultiplay=%s,bIsPvP=%s,bCanPickupOtherGuildDeathPenaltyDrop=%s,bEnableNonLoginPenalty=%s,bEnableFastTravel=%s,bIsStartLocationSelectByMap=%s,bExistPlayerAfterLogout=%s,bEnableDefenseOtherGuildPlayer=%s,CoopPlayerMaxNum=%d,ServerPlayerMaxNum=%d,ServerName=%s,ServerDescription=%s,AdminPassword=%s,ServerPassword=%s,PublicPort=%d,PublicIP=%s,RCONEnabled=%s,RCONPort=%d,Region=%s,bUseAuth=%s,BanListURL=%s,RESTAPIEnabled=%s,RESTAPIPort=%d,bShowPlayerList=%s,AllowConnectPlatform=%s,bIsUseBackupSaveData=%s,LogFormatType=%s)",
		s.Difficulty, s.DayTimeSpeedRate, s.NightTimeSpeedRate, s.ExpRate, s.PalCaptureRate, s.PalSpawnNumRate, s.PalDamageRateAttack, s.PalDamageRateDefense, s.PlayerDamageRateAttack, s.PlayerDamageRateDefense, s.PlayerStomachDecreaceRate, s.PlayerStaminaDecreaceRate, s.PlayerAutoHPRegeneRate, s.PlayerAutoHpRegeneRateInSleep, s.PalStomachDecreaceRate, s.PalStaminaDecreaceRate, s.PalAutoHPRegeneRate, s.PalAutoHpRegeneRateInSleep, s.BuildObjectDamageRate, s.BuildObjectDeteriorationDamageRate, s.CollectionDropRate, s.CollectionObjectHpRate, s.CollectionObjectRespawnSpeedRate, s.EnemyDropItemRate, s.DeathPenalty, getBoolString(s.BEnablePlayerToPlayerDamage), getBoolString(s.BEnableFriendlyFire), getBoolString(s.BEnableInvaderEnemy), getBoolString(s.BActiveUNKO), getBoolString(s.BEnableAimAssistPad), getBoolString(s.BEnableAimAssistKeyboard), s.DropItemMaxNum, s.DropItemMaxNum_UNKO, s.BaseCampMaxNum, s.BaseCampWorkerMaxNum, s.DropItemAliveMaxHours, getBoolString(s.BAutoResetGuildNoOnlinePlayers), s.AutoResetGuildTimeNoOnlinePlayers, s.GuildPlayerMaxNum, s.PalEggDefaultHatchingTime, s.WorkSpeedRate, getBoolString(s.BIsMultiplay), getBoolString(s.BIsPvP), getBoolString(s.BCanPickupOtherGuildDeathPenaltyDrop), getBoolString(s.BEnableNonLoginPenalty), getBoolString(s.BEnableFastTravel), getBoolString(s.BIsStartLocationSelectByMap), getBoolString(s.BExistPlayerAfterLogout), getBoolString(s.BEnableDefenseOtherGuildPlayer), s.CoopPlayerMaxNum, s.ServerPlayerMaxNum, s.ServerName, s.ServerDescription, s.AdminPassword, s.ServerPassword, s.PublicPort, s.PublicIP, getBoolString(s.RCONEnabled), s.RCONPort, s.Region, getBoolString(s.BUseAuth), s.BanListURL, getBoolString(s.RESTAPIEnabled), s.RESTAPIPort, getBoolString(s.BShowPlayerList), s.AllowConnectPlatform, getBoolString(s.BIsUseBackupSaveData), s.LogFormatType)
	return str
}

var defaultGameConfig = GameConfig{
	Difficulty: "None", DayTimeSpeedRate: 1.000000, NightTimeSpeedRate: 1.000000, ExpRate: 1.000000, PalCaptureRate: 1.000000,
	PalSpawnNumRate: 1.000000, PalDamageRateAttack: 1.000000, PalDamageRateDefense: 1.000000, PlayerDamageRateAttack: 1.000000,
	PlayerDamageRateDefense: 1.000000, PlayerStomachDecreaceRate: 1.000000, PlayerStaminaDecreaceRate: 1.000000,
	PlayerAutoHPRegeneRate: 1.000000, PlayerAutoHpRegeneRateInSleep: 1.000000, PalStomachDecreaceRate: 1.000000,
	PalStaminaDecreaceRate: 1.000000, PalAutoHPRegeneRate: 1.000000, PalAutoHpRegeneRateInSleep: 1.000000,
	BuildObjectDamageRate: 1.000000, BuildObjectDeteriorationDamageRate: 1.000000, CollectionDropRate: 1.000000,
	CollectionObjectHpRate: 1.000000, CollectionObjectRespawnSpeedRate: 1.000000, EnemyDropItemRate: 1.000000,
	DeathPenalty: "All", BEnablePlayerToPlayerDamage: false, BEnableFriendlyFire: false, BEnableInvaderEnemy: true,
	BActiveUNKO: false, BEnableAimAssistPad: true, BEnableAimAssistKeyboard: false, DropItemMaxNum: 3000,
	DropItemMaxNum_UNKO: 100, BaseCampMaxNum: 128, BaseCampWorkerMaxNum: 15, DropItemAliveMaxHours: 1.000000,
	BAutoResetGuildNoOnlinePlayers: false, AutoResetGuildTimeNoOnlinePlayers: 72.000000, GuildPlayerMaxNum: 20,
	PalEggDefaultHatchingTime: 72.000000, WorkSpeedRate: 1.000000, BIsMultiplay: false, BIsPvP: false,
	BCanPickupOtherGuildDeathPenaltyDrop: false, BEnableNonLoginPenalty: true, BEnableFastTravel: true,
	BIsStartLocationSelectByMap: true, BExistPlayerAfterLogout: false, BEnableDefenseOtherGuildPlayer: false,
	CoopPlayerMaxNum: 4, ServerPlayerMaxNum: 32, ServerName: "Default Palworld Server", ServerDescription: "",
	AdminPassword: "", ServerPassword: "", PublicPort: 8211, PublicIP: "", RCONEnabled: false, RCONPort: 25575, Region: "",
	BUseAuth: true, BanListURL: "https://api.palworldgame.com/api/banlist.txt", RESTAPIEnabled: false, RESTAPIPort: 8212,
	BShowPlayerList: false, AllowConnectPlatform: "Steam", BIsUseBackupSaveData: true, LogFormatType: "Text",
}
var EditGameConfig GameConfig

const (
	GameStart = iota
	GameStop
	GameUpdating
	GameRestarting
	GameNotFound
)

// 用户输入PalServer路径和steamcmd路径 保存到配置后，从配置中读取
const STEAM_CMD_PATH = "C:\\steamcmd"
const GAME_PATH = "C:\\steamcmd\\steamapps\\common\\PalServer"
const GAME_NAME = "PalServer.exe"
const GAME_CMD_NAME = "PalServer-Win64-Shipping-Cmd.exe"
const GAME_TO_CMD_PATH = "\\Pal\\Binaries\\Win64\\"

var GameStatus = GameNotFound

var PalSteamCmdID = "2394010"

// 配置文件地址 暂时设为本地地址
const SERVER_CONFIG_FILE = "..\\DefaultPalWorldSettings.ini"
