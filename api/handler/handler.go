package handler

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	Handle(c *gin.Context)
}
