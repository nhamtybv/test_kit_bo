package utils

type conextId string

const (
	OracleConnectionKey conextId = "oracle_connection"
	SettingTable        conextId = "SETTINGS"
	MainBucket          conextId = "CONFIG_BUCKET"
	DatabaseName        conextId = "config.db"
)
