package main

import (
	"api/api"
	"github.com/rs/zerolog/log"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()
	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failded to initialize Server:")
	}
}
