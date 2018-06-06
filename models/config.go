package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
)

type Config struct {
	Templates []Template

	// registers all http methods that the APIs
	// contain
	RegisteryAPIMethods map[string]string

	Servers []Server
}

func NewConfig(t []Template) *Config {

	config := &Config{
		RegisteryAPIMethods: make(map[string]string),
	}

	config.Templates = append(config.Templates, t...)

	config.RegisteryAPIMethods["Get"] = "GET"
	config.RegisteryAPIMethods["Post"] = "POST"
	config.RegisteryAPIMethods["Patch"] = "PATCH"
	config.RegisteryAPIMethods["Delete"] = "DELETE"

	return config
}

func (c *Config) ScanAPIFiles() error {
	fmt.Println(len(c.Templates))
	for i, template := range c.Templates {
		var files []string
		err := filepath.Walk(template.ControllerPath, func(path string, f os.FileInfo, err error) error {
			// fmt.Printf("%d,%s\n", i, path)
			if !f.IsDir() {
				files = append(files, path)
			}
			return nil
		})

		if err != nil {
			return err
		}

		c.Templates[i].fileList = append(c.Templates[i].fileList, files[1:]...)
	}
	return nil
}

func (c *Config) ParseAPIMethods() error {
	for i, template := range c.Templates {
		if len(template.fileList) == 0 {
			return errors.New("No files found")
		}

		for _, fileName := range template.fileList {
			apiName, _ := getAPIName(fileName)
			var apiMethods []string
			for _, value := range c.RegisteryAPIMethods {
				apiMethods = append(apiMethods, value)
			}

			c.Templates[i].APIs[apiName] = append(c.Templates[i].APIs[apiName], apiMethods...)
		}
	}

	return nil
}

func getAPIName(path string) (string, error) {
	re := regexp.MustCompile(`[\w-]+\.`)
	apiName := re.FindString(path)
	apiName = apiName[:len(apiName)-1]

	words := strings.Split(apiName, "_")
	count := len(words)
	for i, word := range words {
		if i == (count - 1) {
			words[i] = inflection.Plural(word)
		}

		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, ""), nil
}

func parseFuncName(line string, funcName map[string]string) (string, error) {
	method := ""
	reg1 := regexp.MustCompile("func").MatchString
	if reg1(line) {
		for key, value := range funcName {
			reg2 := regexp.MustCompile(key).MatchString
			if reg2(line) {
				method = value
			}
		}
	}

	return method, nil
}

func (c *Config) OutputConfigFile(fileName string) error {
	err := c.ScanAPIFiles()
	if err != nil {
		return err
	}

	err = c.ParseAPIMethods()
	if err != nil {
		return err
	}
	for _, template := range c.Templates {

		for apiName, methods := range template.APIs {

			for _, m := range methods {
				path := ""
				method := m
				if template.ServerTemplate.ProjectVersion != "" {
					path = "/" + template.ServerTemplate.ProjectVersion + "/" + apiName
				} else {
					path = "/" + apiName
				}

				if method == "GETBYID" {
					method = "GET"
					path += "/:id"
				}

				s := Server{
					Method:        method,
					Path:          path,
					ProxyScheme:   template.ServerTemplate.ProxySchema,
					ProxyPass:     template.ServerTemplate.ProxyPass,
					ProxyPassPath: "/" + apiName,
					APIVersion:    template.ServerTemplate.APIVersion,
					CustomConfigs: template.ServerTemplate.CustomConfigs,
				}

				c.Servers = append(c.Servers, s)
			}
		}
	}

	err = CreateFile(fileName, c.Servers)
	if err != nil {
		return err
	}

	return nil
}

func CreateFile(filename string, servers []Server) error {
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

	data, err := json.MarshalIndent(servers, "", "	")
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
