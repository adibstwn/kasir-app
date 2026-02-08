package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	Port         string `json:"PORT"`
	DbConnection string `json:"DATABASE_URL"`
	DbGormHost   string `json:"DATABASE_GORM_HOST"`
	DbGormUser   string `json:"DATABASE_GORM_USER"`
	DbGormPass   string `json:"DATABASE_GORM_PASS"`
	DbGormDB     string `json:"DATABASE_GORM_DB"`
	DbGormPort   string `json:"DATABASE_GORM_PORT"`

	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
}

func LoadConfig(path string) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := &Config{
		Port:         viper.GetString("PORT"),
		DbConnection: viper.GetString("DATABASE_URL"),
		DbGormHost:   viper.GetString("DATABASE_GORM_HOST"),
		DbGormUser:   viper.GetString("DATABASE_GORM_USER"),
		DbGormPass:   viper.GetString("DATABASE_GORM_PASS"),
		DbGormDB:     viper.GetString("DATABASE_GORM_DB"),
		DbGormPort:   viper.GetString("DATABASE_GORM_PORT"),
	}

	AppConfig = config
}
