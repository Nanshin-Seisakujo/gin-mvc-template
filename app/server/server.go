package server

import (
	"gin-mvc/app/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
