package config

type ConfigDB struct {
	DbConnection string
	DbHost       string
	DbPort       string
	DbName       string
	DbUserName   string
	DbPassword   string
}

type ConfigBitrix struct {
	BitrixSite  string
	BitrixToken string
}

func SetUp() (bitrix ConfigBitrix, db ConfigDB) {

	bitrix.BitrixSite = "https://intranet.idex.cc/rest/1/"
	bitrix.BitrixToken = "evcwp69f5yg7gkwc"

	db.DbConnection = "mysql"
	db.DbHost = "localhost"
	db.DbPort = "3306"
	db.DbName = "guardias"
	db.DbUserName = "root"
	db.DbPassword = "diegoa91"

	return
}
