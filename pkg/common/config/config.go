package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	DbHost   string `mapstructure:"DB_HOST"`
	DbUser   string `mapstructure:"DB_USER"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"DB_NAME"`
	DbPort   string `mapstructure:"DB_PORT"`
	SslMode  string `mapstructure:"sslmode"`
	TimeZone string `mapstructure:"TIMEZONE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.SetConfigType("env")

	// Viper should read values from .env file and override configs
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
