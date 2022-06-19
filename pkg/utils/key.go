package utils

type conextId string

const (
	OracleConnectionKey conextId = "oracle_connection"
	SettingTable        conextId = "SETTINGS"
	ProductTable        conextId = "PRODUCTS"
	MainBucket          conextId = "CONFIG_BUCKET"
	DatabaseName        conextId = "config.db"
)
