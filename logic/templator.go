package logic

import (
	"errors"
	"reflect"
)

// Templator is a template interface
type Templator interface {
	MapOut() (Mapper, error)
}

// DefaultTemplate stores GateCloud server model
type DefaultTemplate struct {
	Method        string `json:"method"`
	Resource      string `json:"resource"`
	Version       string `json:"version"`
	Path          string `cj:"-"`
	ProxyScheme   string `json:"proxy_scheme"`
	ProxyDomain   string `json:"proxy_domain"`
	ProxyPort     string `json:"proxy_port"`
	ProxyPass     string `cj:"-"`
	ProxyPassPath string `cj:"-"`
	ProxyVersion  string `json:"proxy_version"`
	// CustomConfigs map[string]interface{}
	Pattern Pattern `cj:"-" json:"pattern"`
}

// Default creates a default template instance
func Default() (Templator, error) {
	return &DefaultTemplate{}, nil
}

// MapOut generate mapping instance based on the template declaration
func (t *DefaultTemplate) MapOut() (Mapper, error) {
	m := make(Pattern)
	tStruct := reflect.ValueOf(t).Elem()
	if !tStruct.IsValid() {
		return nil, errors.New("reflect template failure")
	}

	if tStruct.Kind() == reflect.Struct {
		for i := 0; i < tStruct.NumField(); i++ {
			tag := tStruct.Type().Field(i).Tag
			if tag.Get("cj") != "-" {
				name := tStruct.Type().Field(i).Name
				name = "$" + name
				tStruct.Field(i).SetString(name)
				m[name] = []string{}
			}
		}
	}

	return &m, nil
}
