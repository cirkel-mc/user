package main

import (
	"cirkel/user/cmd"

	"github.com/cirkel-mc/goutils/config"
	_ "github.com/joho/godotenv/autoload"
)

const serviceName = "user"

func main() {
	cfg := config.New(serviceName)
	defer cfg.Exit()

	srv := cmd.Serve(cfg)
	srv.Run()
}
