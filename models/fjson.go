package models

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
)

type ConfigParamters map[string]interface{}

type ConfigKeys []string

// New initialises an instance
func New() (ConfigParamters, error) {
	return ConfigParamters(make(map[string]interface{})), nil
}

// Add adds the key and value
func (cp *ConfigParamters) Add(key string, value interface{}) error {
	if len(*cp) == 0 {
		*cp = ConfigParamters(make(map[string]interface{}))
	}
	(*cp)[key] = value
	return nil
}

func (cp *ConfigParamters) CreateFile(filename string) error {
	var (
		file *os.File
		path string
	)

	if filename == "" {
		return errors.New("File name can't be null")
	}

	folderpath, err := os.Getwd()
	if err != nil {
		return err
	}
	path = folderpath + "\\files\\"

	// detect if file exists
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir("files", 0755)
		if err != nil {
			return err
		}
	}

	match := regexp.MustCompile(`[.\d]*.json\z`).MatchString
	if !match(filename) {
		path = path + filename + ".json"
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
