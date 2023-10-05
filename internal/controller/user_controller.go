package controller

import (
	"github.com/3lur/go-mall/internal/service"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userSrv service.UserService
}

// Register implements UserController.
func (*userController) Register(ctx *gin.Context) {
	ctx.String(200, "user register")
}

type UserController interface {
	Register(ctx *gin.Context)
}

func NewUserController(userSrv service.UserService) UserController {
	return &userController{userSrv}
}
