package main

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp/v3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database string `yml:"database"`
}

func main() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	pp.Println(config)

}
