package configs

type redisConfigs struct {
	Uri      string
	Password string
	Database int
	Protocol int
}

type Configuration struct {
	Port          string
	EnvName       string
	RedisConfigs  redisConfigs
	AuthSecretKey []byte
}

var ServiceConfigs Configuration

func LoadServiceConfigurations(filename string) {

	switch configs := filename; configs {

	case "dev":
		conf := Configuration{}
		devConfig := DevConfig{&conf}
		devConfig.ConfigManager()
		ServiceConfigs = *devConfig.Configuration

	case "env":
		conf := Configuration{}
		prodConfig := EnvConfig{&conf}
		prodConfig.ConfigManager()
		ServiceConfigs = *prodConfig.Configuration

	default:
		conf := Configuration{}
		localConfig := DevConfig{&conf}
		localConfig.ConfigManager()
		ServiceConfigs = *localConfig.Configuration
	}

}
