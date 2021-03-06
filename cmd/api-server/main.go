package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"flag"
	"github.com/kantegory/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "conf/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}
	
	s := apiserver.Start(config)
	if err := s; err != nil {
		log.Fatal(err)
	}
}
