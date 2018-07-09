package models

import "github.com/gin-gonic/gin"

type JSONTemplate struct {
	Title   string
	Content string
}

func (jt *JSONTemplate) SaveHandler(ctx *gin.Context) {
	// m := make(map[string]interface{})

}
