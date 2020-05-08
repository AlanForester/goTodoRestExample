package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Postgres struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Host string `json:"host"`
		Name string `json:"name"`
	} `json:"postgres"`
}

func readConfigFile() []byte {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Panicf("Config read error: %s", err.Error())
	}
	return yamlFile
}

func NewConfig() *Config {
	conf := new(Config)
	file := readConfigFile()
	err := yaml.Unmarshal(file, conf)
	if err != nil {
		log.Panicf("Error parse yaml file: %s", err.Error())
	}
	return conf
}
