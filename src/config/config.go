package config

func LoadConfig() {
	LoadEnv()
	ConnectDb()
}
