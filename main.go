package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vahidlotfi71/ticket/internal/util"
)

func main() {
	config , err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("can not load config")
	}

	if config.APPDEBUG == "true" {
		log.logger = log.Output(zerolog.ConsoleWriter{out : os.Stderr})
	}

	if config.APPNAME
}