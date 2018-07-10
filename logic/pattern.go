package logic

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Pattern is alias of data type map[string][]string
type Pattern map[string][]string

// ToKeys extracts all keys of map
func (p *Pattern) ToKeys() ([]string, error) {
	var ks []string
	pStruct := reflect.ValueOf(p).Elem()
	if !pStruct.IsValid() {
		return nil, errors.New("reflect pattern failure")
	}

	if pStruct.Kind() == reflect.Map {
		keys := pStruct.MapKeys()
		for _, key := range keys {
			k, ok := key.Interface().(string)
			if ok {
				ks = append(ks, k)
			}
		}
	}
	return ks, nil
}

// Assemble splits the pattern string to map[string][]string
func (p *Pattern) Assemble(content string) error {

	content = `$Method:GET,POST,PATCH,DELETE;$Resource:Schools,Favourites;
	$Version:v1,v0;$ProxyScheme:http;$ProxyDomain:127.0.0.1;
	$ProxyPort:8021;$ProxyVersion:v1`

	s := strings.Replace(content, "$", "", -1)
	s = strings.Replace(s, "\r\n", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	keyPairs := strings.Split(s, ";")
	for _, v := range keyPairs {
		pairs := strings.Split(strings.TrimSpace(v), ":")
		if len(pairs) != 2 {
			return errors.New("format is invalid")
		}

		values := strings.Split(strings.TrimSpace(pairs[1]), ",")
		if len(values) < 1 {
			return errors.New("pattern value format is invalid")
		}
		(*p)[pairs[0]] = values
	}

	keys, err := p.ToKeys()
	if err != nil {
		return err
	}

	for _, key := range keys {
		p.Export(key)
	}

	fmt.Println(*p)
	return nil
}

// Export exports JSON file
func (p *Pattern) Export(key string) error {
	// values, ok := (*p)[key]
	// if !ok {
	// 	return errors.New("key:" + key + " is not found")
	// }

	// for _, v:=range values{
	// 	p.
	// }
	return nil
}
