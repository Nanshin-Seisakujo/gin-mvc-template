package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Assets struct {
	Css string `json:"css.css"`
	Js  string `json:"main.js"`
}

var config *viper.Viper
var assets Assets

func Init(env string) {
	var err error

	config = viper.New()
	config.SetConfigType("yaml")
	println(env)
	config.SetConfigName(env)
	config.AddConfigPath("../app/config/")
	config.AddConfigPath("app/config/")

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile("./.data/manifest.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &assets); err != nil {
		log.Fatal(err)
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetAssets() Assets {
	return assets
}
