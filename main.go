package main

import (
	"customized-json/models"
)

func main() {
	controllerPath := `C:\Users\project\cms\controllers`

	c := models.NewConfig()
	err := c.OutputConfigFile(controllerPath)
	if err != nil {
		panic(err)
	}

}
