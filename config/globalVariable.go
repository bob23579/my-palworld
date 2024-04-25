package config

type SystemInfo struct {
	MemoryList string
	CpuList    string
	DistList   string
	Time       string
}
type GameConfig struct {
	Difficulty         string
	DayTimeSpeedRate   float32
	NightTimeSpeedRate float32
	ExpRate            float32
	PalCaptureRate     float32
	//Difficulty=None,DayTimeSpeedRate=1.000000,NightTimeSpeedRate=1.000000,ExpRate=1.000000,PalCaptureRate=1.000000,
	PalSpawnNumRate        float32
	PalDamageRateAttack    float32
	PalDamageRateDefense   float32
	PlayerDamageRateAttack float32
	//PalSpawnNumRate=1.000000,PalDamageRateAttack=1.000000,PalDamageRateDefense=1.000000,PlayerDamageRateAttack=1.000000,
	PlayerDamageRateDefense   float32
	PlayerStomachDecreaceRate float32
	PlayerStaminaDecreaceRate float32
	//PlayerDamageRateDefense=1.000000,PlayerStomachDecreaceRate=1.000000,PlayerStaminaDecreaceRate=1.000000,
	PlayerAutoHPRegeneRate        float32
	PlayerAutoHpRegeneRateInSleep float32
	PalStomachDecreaceRate        float32
	PalStaminaDecreaceRate        float32
	PalAutoHPRegeneRate           float32
	PalAutoHpRegeneRateInSleep    float32
	//PlayerAutoHPRegeneRate=1.000000,PlayerAutoHpRegeneRateInSleep=1.000000,PalStomachDecreaceRate=1.000000,
	//PalStaminaDecreaceRate=1.000000,PalAutoHPRegeneRate=1.000000,PalAutoHpRegeneRateInSleep=1.000000,
	BuildObjectDamageRate              float32
	BuildObjectDeteriorationDamageRate float32
	CollectionDropRate                 float32
	CollectionObjectHpRate             float32
	CollectionObjectRespawnSpeedRate   float32
	EnemyDropItemRate                  float32
	//BuildObjectDamageRate=1.000000,BuildObjectDeteriorationDamageRate=1.000000,CollectionDropRate=1.000000,
	//CollectionObjectHpRate=1.000000,CollectionObjectRespawnSpeedRate=1.000000,EnemyDropItemRate=1.000000,
	DeathPenalty                string
	bEnablePlayerToPlayerDamage bool
	bEnableFriendlyFire         bool
	bEnableInvaderEnemy         bool
	bActiveUNKO                 bool
	bEnableAimAssistPad         bool
	bEnableAimAssistKeyboard    bool
	DropItemMaxNum              int
	//DeathPenalty=All,bEnablePlayerToPlayerDamage=False,bEnableFriendlyFire=False,bEnableInvaderEnemy=True,
	//bActiveUNKO=False,bEnableAimAssistPad=True,bEnableAimAssistKeyboard=False,DropItemMaxNum=3000,
	DropItemMaxNum_UNKO               int
	BaseCampMaxNum                    int
	BaseCampWorkerMaxNum              int
	DropItemAliveMaxHours             float32
	bAutoResetGuildNoOnlinePlayers    bool
	AutoResetGuildTimeNoOnlinePlayers float64
	GuildPlayerMaxNum                 int
	//DropItemMaxNum_UNKO=100,BaseCampMaxNum=128,BaseCampWorkerMaxNum=15,DropItemAliveMaxHours=1.000000,
	//bAutoResetGuildNoOnlinePlayers=False,AutoResetGuildTimeNoOnlinePlayers=72.000000,GuildPlayerMaxNum=20,
	PalEggDefaultHatchingTime            float64
	WorkSpeedRate                        float32
	bIsMultiplay                         bool
	bIsPvP                               bool
	bCanPickupOtherGuildDeathPenaltyDrop bool
	bEnableNonLoginPenalty               bool
	bEnableFastTravel                    bool
	//PalEggDefaultHatchingTime=72.000000,WorkSpeedRate=1.000000,bIsMultiplay=False,bIsPvP=False,
	//bCanPickupOtherGuildDeathPenaltyDrop=False,bEnableNonLoginPenalty=True,bEnableFastTravel=True,
	bIsStartLocationSelectByMap    bool
	bExistPlayerAfterLogout        bool
	bEnableDefenseOtherGuildPlayer bool
	//bIsStartLocationSelectByMap=True,bExistPlayerAfterLogout=False,bEnableDefenseOtherGuildPlayer=False,
	CoopPlayerMaxNum   int
	ServerPlayerMaxNum int
	ServerName         string
	ServerDescription  string
	//CoopPlayerMaxNum=4,ServerPlayerMaxNum=32,ServerName="Default Palworld Server",ServerDescription="",
	AdminPassword  string
	ServerPassword string
	PublicPort     int
	PublicIP       string
	RCONEnabled    bool
	RCONPort       int
	Region         string
	//AdminPassword="",ServerPassword="",PublicPort=8211,PublicIP="",RCONEnabled=False,RCONPort=25575,Region="",
	bUseAuth             bool
	BanListURL           string
	RESTAPIEnabled       bool
	RESTAPIPort          int
	bShowPlayerList      bool
	AllowConnectPlatform string
	bIsUseBackupSaveData bool
	LogFormatType        string
	//bUseAuth=True,BanListURL="https://api.palworldgame.com/api/banlist.txt",RESTAPIEnabled=False,RESTAPIPort=8212,
	//bShowPlayerList=False,AllowConnectPlatform=Steam,bIsUseBackupSaveData=True,LogFormatType=Text)
}

var defaultGameConfig = GameConfig{
	Difficulty: "None", DayTimeSpeedRate: 1.000000, NightTimeSpeedRate: 1.000000, ExpRate: 1.000000, PalCaptureRate: 1.000000,
	PalSpawnNumRate: 1.000000, PalDamageRateAttack: 1.000000, PalDamageRateDefense: 1.000000, PlayerDamageRateAttack: 1.000000,
	PlayerDamageRateDefense: 1.000000, PlayerStomachDecreaceRate: 1.000000, PlayerStaminaDecreaceRate: 1.000000,
	PlayerAutoHPRegeneRate: 1.000000, PlayerAutoHpRegeneRateInSleep: 1.000000, PalStomachDecreaceRate: 1.000000,
	PalStaminaDecreaceRate: 1.000000, PalAutoHPRegeneRate: 1.000000, PalAutoHpRegeneRateInSleep: 1.000000,
	BuildObjectDamageRate: 1.000000, BuildObjectDeteriorationDamageRate: 1.000000, CollectionDropRate: 1.000000,
	CollectionObjectHpRate: 1.000000, CollectionObjectRespawnSpeedRate: 1.000000, EnemyDropItemRate: 1.000000,
	DeathPenalty: "All", bEnablePlayerToPlayerDamage: false, bEnableFriendlyFire: false, bEnableInvaderEnemy: true,
	bActiveUNKO: false, bEnableAimAssistPad: true, bEnableAimAssistKeyboard: false, DropItemMaxNum: 3000,
	DropItemMaxNum_UNKO: 100, BaseCampMaxNum: 128, BaseCampWorkerMaxNum: 15, DropItemAliveMaxHours: 1.000000,
	bAutoResetGuildNoOnlinePlayers: false, AutoResetGuildTimeNoOnlinePlayers: 72.000000, GuildPlayerMaxNum: 20,
	PalEggDefaultHatchingTime: 72.000000, WorkSpeedRate: 1.000000, bIsMultiplay: false, bIsPvP: false,
	bCanPickupOtherGuildDeathPenaltyDrop: false, bEnableNonLoginPenalty: true, bEnableFastTravel: true,
	bIsStartLocationSelectByMap: true, bExistPlayerAfterLogout: false, bEnableDefenseOtherGuildPlayer: false,
	CoopPlayerMaxNum: 4, ServerPlayerMaxNum: 32, ServerName: "Default Palworld Server", ServerDescription: "",
	AdminPassword: "", ServerPassword: "", PublicPort: 8211, PublicIP: "", RCONEnabled: false, RCONPort: 25575, Region: "",
	bUseAuth: true, BanListURL: "https://api.palworldgame.com/api/banlist.txt", RESTAPIEnabled: false, RESTAPIPort: 8212,
	bShowPlayerList: false, AllowConnectPlatform: "Steam", bIsUseBackupSaveData: true, LogFormatType: "Text",
}
var editGameConfig GameConfig

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
