package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vahidlotfi71/ticket/internal/util"
)

func main() {
	// تنظیمات کانفیگ
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Connot load Config failed")
	}

	if config.APPDEBUG == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// ارتباط با دیتا بیس

	dsn := fmt.Sprintf("postgres;//%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUSERNAME,
		config.DBPASSWORD,
		config.DBHOST,
		config.DBPORT,
		config.DBDATABASE,
	)
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot to conncte database")
	}
	defer dbpool.Close()

	

}
