package main

import (
	"customized-json/models"
	"fmt"
)

func main() {
	cp, err := New()
	if err != nil {
		panic(err)
	}

	cp.Add("Method", "GET")
	fmt.Println(cp)

	err = cp.CreateFile("ttt")
	if err != nil {
		panic(err)
	}
}

// New initialises an instance
func New() (models.ConfigParamters, error) {
	return models.ConfigParamters(make(map[string]interface{})), nil
}
