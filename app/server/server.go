package server

import (
	"kleiberman-app/app/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
