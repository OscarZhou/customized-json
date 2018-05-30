package models

import (
	"os"
	"path/filepath"
)

type Config struct {
	Templates []ServerTemplate
	FileList  []string
}

func (c *Config) scanAPIFiles(path string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		c.FileList = append(c.FileList, path)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
