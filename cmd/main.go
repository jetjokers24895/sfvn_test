package main

import (
	application "sfvn_test"
	"sfvn_test/config"
)

// @BasePath /
func main() {
	cfg := config.NewConfig()

	application.Start(cfg)
}
