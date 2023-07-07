package main

import (
	"log"

	"github.com/gitaepark/codedeploy-go/server/config"
	"github.com/gitaepark/codedeploy-go/server/loader"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	runGinServer(config)
}

func runGinServer(config config.Config) {
	server, err := loader.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}