package env

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	_ = godotenv.Load()
	viper.AutomaticEnv()
}

func Get(key string, defaultValue ...string) string {
	if v := viper.GetString(key); v != "" {
		return v
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func IsDev() bool { return Get("APP_ENV") == "development" }
