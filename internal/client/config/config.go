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
		HostnameToVerify string `yaml:"hostnameToVerify"`
		CertificateAuthorities string `yaml:"certificateAuthorities"`
	}
	Client struct {
		Certificate string
		Key string
	}
}

var Config _Config

func init() {

	configurationFile := os.Getenv("ORIHIME_CONFIG")
	if len(configurationFile) == 0 {
		log.Fatal("ORIHIME_CONFIG is not defined")
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
