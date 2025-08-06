package main

import (
	"flag"

	"github.com/lyyyuna/storage-manager/pkg/config"
	"github.com/lyyyuna/storage-manager/pkg/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config.yaml", "specify the config file path")
}

func main() {
	cfg := config.NewConfig(configPath)

	s := server.NewServer(cfg)

	s.Run()
}
