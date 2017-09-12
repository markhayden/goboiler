package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	LOCALCONFIG = "local.json"
)

var (
	Conf Configuration
)

type Configuration struct {
	Keys struct {
		A string `json:"a"`
		B string `json:"b"`
	} `json:"keys"`
}

func init() {
	path := "../lib/config/" + LOCALCONFIG

	// as an example we always load config but additional logic can live here to select based on env
	config := LOCALCONFIG

	// get the config yaml file
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("unable to load config <%s>: %v ", config, err)
	}

	log.Printf("config: %s", configFile)

	// unmarshal the yaml
	c := Configuration{}
	err = json.Unmarshal(configFile, &c)
	if err != nil {
		log.Fatalf("unable to unmarshal config <%s>: %v", config, err)
	}

	// set the top level config
	Conf = c

	return
}
