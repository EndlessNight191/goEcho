package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func CreateConnection() error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.sslmode"),
	)

	if DB, err = sql.Open("postgres", dsn); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = DB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
