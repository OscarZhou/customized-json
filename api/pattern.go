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

	// ctx.JSON(http.StatusOK, keys)
	ctx.HTML(http.StatusOK, "add_pattern.html", gin.H{
		"Pattern": keys,
	})
	return
}

// SetPattern sets the pattern's values
func SetPattern(ctx *gin.Context, p logic.Assembler) {
	s := ctx.PostForm("content")
	err := p.Assemble(s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
}
