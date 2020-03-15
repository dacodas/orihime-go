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
	}
}

var Config _Config

func init() {

	contents, err := ioutil.ReadFile(os.Getenv("ORIHIME_SERVER_CONFIG"))
	if err != nil {
		panic(err)
	}
	log.Printf("File: %v", string(contents))

	err = yaml.Unmarshal(contents, &Config)
	if err != nil {
		panic(err)
	}

	log.Printf("Using config: %v", Config)
}
