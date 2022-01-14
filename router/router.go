package router

import (
	"github.com/gin-gonic/gin"
	"go/foundation/restful/handler/status"
	"go/foundation/restful/router/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 检查handlers
	svcd := g.Group("/status")
	{
		svcd.GET("/status", status.ShowStatus)
		svcd.GET("/disk", status.ShowDisk)
		svcd.GET("/mem", status.ShowMEM)
	}
	return g
}
