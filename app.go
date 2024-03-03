package app

import (
	"log"
	"sfvn_test/config"
	"sfvn_test/server"
	"sfvn_test/server/routes"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port) //Todo add config env
	if err != nil {
		log.Fatal("Port already used")
	}
}
