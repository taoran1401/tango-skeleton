package controller

import (
	"github.com/gin-gonic/gin"
)

// index controller
type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (this *IndexController) Index(ctx *gin.Context) {
}
