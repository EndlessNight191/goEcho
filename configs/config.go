package configs

import (
	"github.com/spf13/viper"
)

var cfg *viper.Viper

func GetConfig() error {
	viper.AddConfigPath("../")
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
