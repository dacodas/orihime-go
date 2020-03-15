package config

import (
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type _Config struct {
	Server struct {
		Endpoint string
	}
}

var Config _Config

func init() {

	contents, err := ioutil.ReadFile(os.Getenv("ORIHIME_CONFIG"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(contents, &Config)
	if err != nil {
		panic(err)
	}

	log.Printf("Using config: %v", Config)
}
