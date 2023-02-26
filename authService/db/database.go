package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func CreateConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.sslmode"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping: %v", err)
	}

	return db, nil
}
