package global

type ConfigData struct {
	Db struct {
		Dbname   string
		Username string
		Password string
	}
}

var Config *ConfigData
