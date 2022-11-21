package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Add(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
