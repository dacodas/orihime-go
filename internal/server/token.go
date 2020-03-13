package server

import (
	"encoding/json"
)

type Conf struct {
	User string
	Expires string
}

var Configuration Conf

func ReadConfiguration(_json []byte) error {
	err := json.Unmarshal(_json, &Configuration)
	return err
}

func init() {

}
