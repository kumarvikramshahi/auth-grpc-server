package configs

type DevConfig struct {
	*Configuration
}

func (conf *DevConfig) ConfigManager() {
	conf.Port = "8000"
	conf.EnvName = "dev"
	conf.AuthSecretKey = []byte("Mitran Di Billo")
	conf.RedisConfigs.Uri = "localhost:6379"
	conf.RedisConfigs.Database = 0
	conf.RedisConfigs.Password = ""
	conf.RedisConfigs.Protocol = 2
}
