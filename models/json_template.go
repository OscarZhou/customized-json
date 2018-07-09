package models

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

// JSONTemplate is the file model used to store the structures
type JSONTemplate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Save save the current types to .txt file
func (jt *JSONTemplate) Save() error {
	if jt.Title == "" || jt.Content == "" {
		return errors.New("title or content should not be empty")
	}
	jt.Content = strings.Replace(jt.Content, ",", ",\n", -1)
	return ioutil.WriteFile("struct_"+jt.Title+".txt", []byte(jt.Content), 0600)
}

// SaveHandler
func (jt *JSONTemplate) SaveHandler(ctx *gin.Context) {
	// m := make(map[string]interface{})

}
