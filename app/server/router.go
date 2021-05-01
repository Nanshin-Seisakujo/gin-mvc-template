package server

import (
	"encoding/json"
	"io/ioutil"
	"kleiberman-app/app/controllers"
	"kleiberman-app/app/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Assets struct {
	Css string `json:"css.css"`
	Js  string `json:"main.js"`
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Delims("<%", "%>")
	router.LoadHTMLGlob("app/views/**/*.tmpl")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	bytes, err := ioutil.ReadFile("./.data/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	var assets Assets
	if err := json.Unmarshal(bytes, &assets); err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		js := assets.Js
		css := assets.Css
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "Main Page",
			"main_js":  js,
			"main_css": css,
		})
	})

	// call health controller
	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())

	return router
}
