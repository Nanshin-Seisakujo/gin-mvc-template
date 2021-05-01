package middlewares

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserId interface{}
}

var Session SessionInfo

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		Session.UserId = session.Get("UserId")

		if Session.UserId == nil {
			log.Println("You are not logged in.")
			c.Next()
		} else {
			c.Set("UserId", Session.UserId)
			c.Next()
		}
	}
}
