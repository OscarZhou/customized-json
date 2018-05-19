package models

import (
	"encoding/json"
	"errors"
	"os"
)

var path = "files/"

type ConfigParamters map[string]interface{}

type ConfigKeys []string

// Add adds the key and value
func (cp *ConfigParamters) Add(key string, value interface{}) error {
	if len(*cp) == 0 {
		*cp = ConfigParamters(make(map[string]interface{}))
	}
	(*cp)[key] = value
	return nil
}

func (cp *ConfigParamters) CreateFile(filename string) error {
	if filename == "" {
		return errors.New("File name can't be null")
	}
	var file *os.File

	filename = path + filename
	// detect if file exists
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}

	// create file if not exists
	if !os.IsNotExist(err) {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}

	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	data, err := json.Marshal(*cp)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}
