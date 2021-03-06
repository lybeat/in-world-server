package net

import (
	"github.com/lybeat/in-world-server/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code":    errCode,
		"message": e.GetMsg(errCode),
		"data":    data,
	})

	return
}
