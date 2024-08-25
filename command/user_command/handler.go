package usercommand

import (
	"github.com/gin-gonic/gin"
)

type (
	Handler interface {
		Create(ctx *gin.Context)
	}
	handler struct {
		usecase Usecase
	}
	params struct {
		Name string `json:"name"`
	}
)

func (h *handler) Create(ctx *gin.Context) {
	var p params
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	args := &createArgs{
		name: p.Name,
	}
	if err := h.usecase.Create(args); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User created"})
}
