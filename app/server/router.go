package server

import (
	"gin-mvc/app/config"
	"gin-mvc/app/controllers"
	"gin-mvc/app/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Delims("<%", "%>")
	router.LoadHTMLGlob("app/views/**/*.html")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	assets := config.GetAssets()

	router.GET("/", func(c *gin.Context) {
		js := assets.Js
		css := assets.Css
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Main Page",
			"main_js":  js,
			"main_css": css,
		})
	})

	// call health controller
	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	// auth middleware
	router.Use(middlewares.AuthMiddleware())

	// cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Origin"},
		AllowCredentials: true,
	}))

	return router
}
