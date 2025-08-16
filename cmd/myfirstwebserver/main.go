package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/ttahaiyana/my-first-web-server/internal/app/myfirstwebserver"
)

var (
	configPath *string
)

func init() {
	configPath = flag.String("path", "configs/myfirstwebserver.toml", "path to configuration file")
}

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Println("faild read from file. using default configs. ", err)
	}

	config := myfirstwebserver.NewConfig()

	_, err = toml.Decode(string(data), &config)
	if err != nil {
		log.Println("faild read from file. using default configs. ", err)
	}

	server := myfirstwebserver.New(*config)

	log.Fatal(server.Start())
}
