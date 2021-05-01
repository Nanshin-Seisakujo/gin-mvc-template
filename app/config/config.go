package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

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
