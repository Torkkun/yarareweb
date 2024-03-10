package handler

import "github.com/gin-gonic/gin"

func xsshandler(ctx *gin.Context) {
	ctx.HTML(200, "xss.html", nil)
}
