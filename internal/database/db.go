package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func CreateConnection() error {
	var err error

	if DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.sslmode"))); err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}
