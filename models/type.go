package models

import (
	"bufio"
	"fmt"
	"os"
)

var basicTypes = []string{
	"String",
	"Numeric",
	"Bool",
}

func Append(newItem string) {
	basicTypes = append(basicTypes, newItem)
}

func GetAll() []string {
	return basicTypes
}

func Save() error {
	file, err := os.Create("type.txt")
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range basicTypes {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
