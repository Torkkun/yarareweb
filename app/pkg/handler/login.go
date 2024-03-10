package handler

import (
	"OCsemmerApp/pkg/domain"
	"OCsemmerApp/pkg/model"
	"log"

	"github.com/gin-gonic/gin"
)

// GET /login
func loginHandler(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil)
}

// POST /autenticate
func autenticateHandler(ctx *gin.Context) {
	login := domain.LoginRequest{}
	if err := ctx.Bind(&login); err != nil {
		log.Println(err)
		ctx.JSON(400, "Bad Request")
		return
	}
	log.Println(login)
	if login.Name == "" && login.Pass == "" {
		log.Println("nil request")
		ctx.JSON(400, "not Request")
		return
	}
	if err := model.SearchUser(login); err != nil {
		log.Println(err)
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, "ログイン成功")
}
