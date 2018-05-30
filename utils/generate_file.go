package utils

import (
	"customized-json/models"
	"encoding/json"
	"errors"
	"os"
	"regexp"
)

func CreateFile(filename string, servers []models.Server) error {
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

	data, err := json.Marshal(servers)
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
