package main

import (
	"cirkel/user/cmd"

	"github.com/cirkel-mc/goutils/config"
)

const serviceName = "user"

func main() {
	cfg := config.New(serviceName)
	defer cfg.Exit()

	srv := cmd.Serve(cfg)
	srv.Run()
}
