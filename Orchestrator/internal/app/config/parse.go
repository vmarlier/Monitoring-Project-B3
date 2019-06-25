package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// This package will be able to parse and read a configuration files for the metrics monitoring
// The configuration file will be on the host file system
// This package will be call by the "server" package to configure himself

// Config is the struct version of the json config file
// you can follow that struct to recreate a json config file
type Config struct {
	Host    string `json:"host"`
	Refresh Refr   `json:"refresh"`
	Disks   []Disk `json:"disks"`
}

type Refr struct {
	Unit string `json:"unit"`
	Time int    `json:"time"`
}

type Disk struct {
	Path string `json:"path"`
}

// ParseConf return an Config type struct
func ParseConf(path string) (conf Config) {
	file, err := os.Open(path)
	handleErr(err)
	if err != nil {
		log.Fatal("WARNING: CONFIGURATIONS FILES MISSING")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	handleErr(err)

	json.Unmarshal(bytes, &conf)

	return conf
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
