package main

import (
	"encoding/json"
	"io/ioutil"
	"kal/Godeps/_workspace/src/bitbucket.org/kardianos/osext"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	NotActive bool   `json:"notActive"`
	RelRoot   string `json:"relRoot"`
}

var config = readConfig()

func readConfig() Config {
	var c Config
	if exep, err := osext.ExecutableFolder(); err == nil {
		dir, _ := filepath.Split(strings.TrimSuffix(exep, string(os.PathSeparator)))
		if file, err := ioutil.ReadFile(filepath.Join(dir, "config")); err == nil {
			if err = json.Unmarshal(file, &c); err == nil {
				return c
			}
		}
	}
	// Ugly, but we have to eat an error here. It is to early to properly report what is going on,
	// so all configuration values should behave properly when not initialized
	return Config{}
}
