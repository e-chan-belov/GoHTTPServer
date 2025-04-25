package main

import (
	"hw2/config"
	_ "hw2/docs"
	"hw2/pkg/factory"
	pkgHttp "hw2/pkg/http"
	"log"
)

// @title My API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
func main() {

	appFlags := config.ParseFlags()
	var cfg config.HTTPConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)
	address := &cfg.Address

	r := factory.NewServer()

	log.Printf("Starting server on %s", *address)
	if err := pkgHttp.CreateAndRunServer(r, *address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
