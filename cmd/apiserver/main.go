package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/larikhide/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string //for send path to config file as a flag
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse() //for parsing flag into variables

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config) //write toml file to variables in config
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config) //starts server
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
