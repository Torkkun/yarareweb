package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		Engine: gin.Default(),
	}
	router.Engine.StaticFS("/static", http.Dir("public"))
	router.Engine.LoadHTMLGlob("templates/*.html")
	// portal page
	router.Engine.GET("/", func(ctx *gin.Context) {
		indexportalHandler(ctx)
	})
	// login test page
	router.Engine.GET("/login", func(ctx *gin.Context) {
		loginHandler(ctx)
	})
	router.Engine.POST("/authenticate", func(ctx *gin.Context) {
		autenticateHandler(ctx)
	})
	// sql injection test page
	datapage := router.Engine.Group("/data")
	{
		datapage.GET("/", func(ctx *gin.Context) {
			datahandler(ctx)
		})
		datapage.GET("/search", func(ctx *gin.Context) {
			searchhandler(ctx)
		})
		datapage.GET("/test", func(ctx *gin.Context) {
			testsearchhandler(ctx)
		})
	}
	router.Engine.GET("/xss", func(ctx *gin.Context) {
		xsshandler(ctx)
	})
	return router
}

func (router *Router) Run() {
	router.Engine.Run(":8080")
}
