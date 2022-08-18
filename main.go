package main

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp/v3"
	"gopkg.in/yaml.v3"
)

var Config = struct {
	Database struct {
		URL string
	}
}{}

func init() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal([]byte(data), &Config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {
	SomeService()
}

func SomeService() {
	pp.Println(Config.Database)
}
