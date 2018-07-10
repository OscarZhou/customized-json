package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
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
	// jt.Content = strings.Replace(jt.Content, ",", ",\n", -1)
	return ioutil.WriteFile("files/struct_"+jt.Title+".txt", []byte(jt.Content), 0600)
}

// Parse parses the .txt file to struct
func (jt *JSONTemplate) Parse() (ModelStruct, error) {
	var modelStruct ModelStruct
	keyPairs := strings.Split(jt.Content, ",")
	modelStruct.ModelTitle = jt.Title
	for _, v := range keyPairs {
		pairs := strings.TrimSpace(v)
		s := strings.Split(pairs, " ")
		if len(s) != 2 {
			return ModelStruct{}, errors.New("format is invalid")
		}
		modelStruct.KeyTypes = append(modelStruct.KeyTypes, KeyType{Key: s[0], Type: s[1]})
	}
	return modelStruct, nil
}

// LoadJSONTemplate loads struct_ file by file name
func LoadJSONTemplate(filename string) (*JSONTemplate, error) {
	match := regexp.MustCompile("struct_").MatchString
	if !match(filename) {
		return nil, errors.New("the file is not model struct file")
	}
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	title := strings.Replace(filename, "files\\", "", -1)
	title = strings.Replace(title, "struct_", "", -1)
	title = strings.Replace(title, ".txt", "", -1)
	fmt.Println("title is ", title)
	fmt.Println("content is", string(body))
	return &JSONTemplate{Title: title, Content: string(body)}, nil
}
