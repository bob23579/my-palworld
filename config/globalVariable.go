package config

import "fmt"

type SystemInfo struct {
	MemoryList string
	CpuList    string
	DistList   string
	Time       string
}
type GameConfig struct {
	Difficulty                           string  `json:"difficulty"`
	DayTimeSpeedRate                     float64 `json:"day_time_speed_rate"`
	NightTimeSpeedRate                   float64 `json:"night_time_speed_rate"`
	ExpRate                              float64 `json:"exp_rate"`
	PalCaptureRate                       float64 `json:"pal_capture_rate"`
	PalSpawnNumRate                      float64 `json:"pal_spawn_num_rate"`
	PalDamageRateAttack                  float64 `json:"pal_damage_rate_attack"`
	PalDamageRateDefense                 float64 `json:"pal_damage_rate_defense"`
	PlayerDamageRateAttack               float64 `json:"player_damage_rate_attack"`
	PlayerDamageRateDefense              float64 `json:"player_damage_rate_defense"`
	PlayerStomachDecreaceRate            float64 `json:"player_stomach_decreace_rate"`
	PlayerStaminaDecreaceRate            float64 `json:"player_stamina_decreace_rate"`
	PlayerAutoHPRegeneRate               float64 `json:"player_auto_hp_regene_rate"`
	PlayerAutoHpRegeneRateInSleep        float64 `json:"player_auto_hp_regene_rate_in_sleep"`
	PalStomachDecreaceRate               float64 `json:"pal_stomach_decreace_rate"`
	PalStaminaDecreaceRate               float64 `json:"pal_stamina_decreace_rate"`
	PalAutoHPRegeneRate                  float64 `json:"pal_auto_hp_regene_rate"`
	PalAutoHpRegeneRateInSleep           float64 `json:"pal_auto_hp_regene_rate_in_sleep"`
	BuildObjectDamageRate                float64 `json:"build_object_damage_rate"`
	BuildObjectDeteriorationDamageRate   float64 `json:"build_object_deterioration_damage_rate"`
	CollectionDropRate                   float64 `json:"collection_drop_rate"`
	CollectionObjectHpRate               float64 `json:"collection_object_hp_rate"`
	CollectionObjectRespawnSpeedRate     float64 `json:"collection_object_respawn_speed_rate"`
	EnemyDropItemRate                    float64 `json:"enemy_drop_item_rate"`
	DeathPenalty                         string  `json:"death_penalty"`
	BEnablePlayerToPlayerDamage          bool    `json:"b_enable_player_to_player_damage"`
	BEnableFriendlyFire                  bool    `json:"b_enable_friendly_fire"`
	BEnableInvaderEnemy                  bool    `json:"b_enable_invader_enemy"`
	BActiveUNKO                          bool    `json:"b_active_unko"`
	BEnableAimAssistPad                  bool    `json:"b_enable_aim_assist_pad"`
	BEnableAimAssistKeyboard             bool    `json:"b_enable_aim_assist_keyboard"`
	DropItemMaxNum                       int     `json:"drop_item_max_num"`
	DropItemMaxNum_UNKO                  int     `json:"drop_item_max_num_unko"`
	BaseCampMaxNum                       int     `json:"base_camp_max_num"`
	BaseCampWorkerMaxNum                 int     `json:"base_camp_worker_max_num"`
	DropItemAliveMaxHours                float64 `json:"drop_item_alive_max_hours"`
	BAutoResetGuildNoOnlinePlayers       bool    `json:"b_auto_reset_guild_no_online_players"`
	AutoResetGuildTimeNoOnlinePlayers    float64 `json:"auto_reset_guild_time_no_online_players"`
	GuildPlayerMaxNum                    int     `json:"guild_player_max_num"`
	PalEggDefaultHatchingTime            float64 `json:"pal_egg_default_hatching_time"`
	WorkSpeedRate                        float64 `json:"work_speed_rate"`
	BIsMultiplay                         bool    `json:"b_is_multiplay"`
	BIsPvP                               bool    `json:"b_is_pv_p"`
	BCanPickupOtherGuildDeathPenaltyDrop bool    `json:"b_can_pickup_other_guild_death_penalty_drop"`
	BEnableNonLoginPenalty               bool    `json:"b_enable_non_login_penalty"`
	BEnableFastTravel                    bool    `json:"b_enable_fast_travel"`
	BIsStartLocationSelectByMap          bool    `json:"b_is_start_location_select_by_map"`
	BExistPlayerAfterLogout              bool    `json:"b_exist_player_after_logout"`
	BEnableDefenseOtherGuildPlayer       bool    `json:"b_enable_defense_other_guild_player"`
	CoopPlayerMaxNum                     int     `json:"coop_player_max_num"`
	ServerPlayerMaxNum                   int     `json:"server_player_max_num"`
	ServerName                           string  `json:"server_name"`
	ServerDescription                    string  `json:"server_description"`
	AdminPassword                        string  `json:"admin_password"`
	ServerPassword                       string  `json:"server_password"`
	PublicPort                           int     `json:"public_port"`
	PublicIP                             string  `json:"public_ip"`
	RCONEnabled                          bool    `json:"rcon_enabled"`
	RCONPort                             int     `json:"rcon_port"`
	Region                               string  `json:"region"`
	BUseAuth                             bool    `json:"b_use_auth"`
	BanListURL                           string  `json:"ban_list_url"`
	RESTAPIEnabled                       bool    `json:"restapi_enabled"`
	RESTAPIPort                          int     `json:"restapi_port"`
	BShowPlayerList                      bool    `json:"b_show_player_list"`
	AllowConnectPlatform                 string  `json:"allow_connect_platform"`
	BIsUseBackupSaveData                 bool    `json:"b_is_use_backup_save_data"`
	LogFormatType                        string  `json:"log_format_type"`
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
	str := fmt.Sprintf("OptionSettings=(Difficulty=%s,DayTimeSpeedRate=%f,NightTimeSpeedRate=%f,ExpRate=%f,PalCaptureRate=%f,PalSpawnNumRate=%f,PalDamageRateAttack=%f,PalDamageRateDefense=%f,PlayerDamageRateAttack=%f,PlayerDamageRateDefense=%f,PlayerStomachDecreaceRate=%f,PlayerStaminaDecreaceRate=%f,PlayerAutoHPRegeneRate=%f,PlayerAutoHpRegeneRateInSleep=%f,PalStomachDecreaceRate=%f,PalStaminaDecreaceRate=%f,PalAutoHPRegeneRate=%f,PalAutoHpRegeneRateInSleep=%f,BuildObjectDamageRate=%f,BuildObjectDeteriorationDamageRate=%f,CollectionDropRate=%f,CollectionObjectHpRate=%f,CollectionObjectRespawnSpeedRate=%f,EnemyDropItemRate=%f,DeathPenalty=%s,bEnablePlayerToPlayerDamage=%s,bEnableFriendlyFire=%s,bEnableInvaderEnemy=%s,bActiveUNKO=%s,bEnableAimAssistPad=%s,bEnableAimAssistKeyboard=%s,DropItemMaxNum=%d,DropItemMaxNum_UNKO=%d,BaseCampMaxNum=%d,BaseCampWorkerMaxNum=%d,DropItemAliveMaxHours=%f,bAutoResetGuildNoOnlinePlayers=%s,AutoResetGuildTimeNoOnlinePlayers=%f,GuildPlayerMaxNum=%d,PalEggDefaultHatchingTime=%f,WorkSpeedRate=%f,bIsMultiplay=%s,bIsPvP=%s,bCanPickupOtherGuildDeathPenaltyDrop=%s,bEnableNonLoginPenalty=%s,bEnableFastTravel=%s,bIsStartLocationSelectByMap=%s,bExistPlayerAfterLogout=%s,bEnableDefenseOtherGuildPlayer=%s,CoopPlayerMaxNum=%d,ServerPlayerMaxNum=%d,ServerName=\"%s\",ServerDescription=\"%s\",AdminPassword=\"%s\",ServerPassword=\"%s\",PublicPort=%d,PublicIP=\"%s\",RCONEnabled=%s,RCONPort=%d,Region=\"%s\",bUseAuth=%s,BanListURL=\"%s\",RESTAPIEnabled=%s,RESTAPIPort=%d,bShowPlayerList=%s,AllowConnectPlatform=%s,bIsUseBackupSaveData=%s,LogFormatType=%s)",
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
