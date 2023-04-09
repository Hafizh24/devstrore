package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBConnection string `mapstructure:"DB_CONNECTION"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
	LogLevel     string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig(fileconfigpath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileconfigpath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.Unmarshal(&config)
	return config, nil
}
