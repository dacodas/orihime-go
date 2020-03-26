package config

import (
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type _Config struct {
	Database struct {
		DataSourceName string `yaml:"dataSourceName"`
	}
	Server struct {
		Listen string
		Key string
		Certificate string
	}
}

var Config _Config

func init() {
	configurationFile := os.Getenv("ORIHIME_SERVER_CONFIG")
	if len(configurationFile) == 0 {
		log.Fatal("ORIHIME_SERVER_CONFIG is not defined")
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(contents, &Config)
	if err != nil {
		panic(err)
	}

	log.Printf("Using config: %v", Config)
}
