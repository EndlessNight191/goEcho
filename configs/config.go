package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port  string
	DbUrl string
}

func GetConfig() Config {
	viper.SetConfigFile("./.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error trying read from config: %w", err)
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	result := Config{
		Port:  port,
		DbUrl: dbUrl,
	}

	return result
}
