package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `json:"PORT"`
	DbConnection string `json:"DATABASE_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := &Config{
		Port:         viper.GetString("PORT"),
		DbConnection: viper.GetString("DATABASE_URL"),
	}
	return config, nil
}
