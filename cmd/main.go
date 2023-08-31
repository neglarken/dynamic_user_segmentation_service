package main

import (
	"flag"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/neglarken/dynamic_user_segmentation_service/config"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/httpserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to config file")
}

// @title           Dynamic User Segmentation Service
// @version         1.0
// @description     This is a Dynamic User Segmentation Service.

// @contact.name   Keril
// @contact.email  khe.14@yandex.ru

// @host      localhost:8080
// @BasePath  /
func main() {
	flag.Parse()

	config := config.NewConfig(configPath)

	err := cleanenv.ReadConfig(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := httpserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
