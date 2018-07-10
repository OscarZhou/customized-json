package api

import (
	"customized-json/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPattern returns the current pattern keys
func GetPattern(ctx *gin.Context, p logic.Mapper) {
	keys, err := p.ToKeys()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, keys)
	return
}
