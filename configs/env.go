package configs

import (
	"os"
	"strconv"
)

type EnvConfig struct {
	*Configuration
}

func (conf *EnvConfig) ConfigManager() {
	conf.Port = os.Getenv("")
	conf.EnvName = os.Getenv(("ENV_NAME"))
	conf.AuthSecretKey = []byte(os.Getenv("AUTH_SECRET_KEY"))
	conf.RedisConfigs.Uri = os.Getenv("REDIS_URI")
	conf.RedisConfigs.Database = convertToInt(os.Getenv("REDIS_DB"))
	conf.RedisConfigs.Password = os.Getenv("REDIS_PASS")
	conf.RedisConfigs.Protocol = convertToInt(os.Getenv("REDIS_PROTOCOL"))
}

func convertToInt(
	value string,
) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}
