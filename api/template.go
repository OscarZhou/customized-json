package api

import (
	"customized-json/logic"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// GetTemplate returns the current pattern keys
func GetTemplate(ctx *gin.Context, t logic.Templator) {

	template, ok := t.(*logic.DefaultTemplate)
	if ok {
		keyTypes := make(map[string]string)
		tRef := reflect.ValueOf(template).Elem()
		if !tRef.IsValid() {
			ctx.JSON(http.StatusInternalServerError, "reflect template failure")
			return
		}

		if tRef.Kind() == reflect.Struct {
			for i := 0; i < tRef.NumField(); i++ {
				k := tRef.Type().Field(i).Name
				ty := tRef.Field(i).Type().Name()
				keyTypes[k] = ty
			}
		}

		ctx.HTML(http.StatusOK, "template.tmpl", gin.H{
			"KeyTypes": keyTypes,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, "assert failure")
	return
}
