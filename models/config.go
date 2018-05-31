package models

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
)

type Config struct {
	Templates []ServerTemplate

	// only can be assigned the value automatically
	fileList []string

	// registers all http methods that the APIs
	// contain
	RegisteryAPIMethods map[string]string

	// key is API name and the value indicates
	// that what methods does this API have
	APIs map[string][]string

	Servers []Server
}

func NewConfig() *Config {

	config := &Config{
		RegisteryAPIMethods: make(map[string]string),
		APIs:                make(map[string][]string),
	}

	customConfig := CustomConfig{
		"Cached":                  true,
		"CachedDurationsInSecond": 10,
		"Authentication":          true,
	}

	serverTemplate := ServerTemplate{
		ProjectVersion: "v0.4",
		APIVersion:     "v1",
		ProxySchema:    "http",
		ProxyPass:      "127.0.0.1:9100",
		CustomConfigs:  customConfig,
	}
	config.Templates = append(config.Templates, serverTemplate)

	config.RegisteryAPIMethods["GetByID"] = "GETBYID"
	config.RegisteryAPIMethods["Get"] = "GET"
	config.RegisteryAPIMethods["Post"] = "POST"
	config.RegisteryAPIMethods["Patch"] = "PATCH"
	config.RegisteryAPIMethods["Delete"] = "DELETE"

	return config
}

func (c *Config) ScanAPIFiles(path string) error {
	var files []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return err
	}

	c.fileList = append(c.fileList, files[1:]...)
	return nil
}

func (c *Config) ParseAPIMethods() error {
	if len(c.fileList) == 0 {
		return errors.New("No files found")
	}

	for _, fileName := range c.fileList {
		apiName, _ := getAPIName(fileName)

		var apiMethods []string
		inFile, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer inFile.Close()

		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			reg1 := regexp.MustCompile("func").MatchString
			if reg1(line) {
				method, err := parseFuncName(line, c.RegisteryAPIMethods)
				if err != nil {
					return err
				}

				if method != "" {
					apiMethods = append(apiMethods, method)
				}
			}
		}

		c.APIs[apiName] = append(c.APIs[apiName], apiMethods...)
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

func (c *Config) OutputConfigFile(controllerPath string) error {
	err := c.ScanAPIFiles(controllerPath)
	if err != nil {
		return err
	}

	err = c.ParseAPIMethods()
	if err != nil {
		return err
	}

	for _, template := range c.Templates {
		for apiName, methods := range c.APIs {
			for _, method := range methods {
				path := "/" + template.APIVersion + "/" + apiName
				if method == "GETBYID" {
					path += "/:id"
				}
				s := Server{
					Method:        method,
					Path:          path,
					ProxyScheme:   template.ProxySchema,
					ProxyPass:     template.ProxyPass,
					ProxyPassPath: "/" + apiName,
					APIVersion:    template.APIVersion,
					CustomConfigs: template.CustomConfigs,
				}

				c.Servers = append(c.Servers, s)
			}
		}
	}

	err = CreateFile("route", c.Servers)
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
