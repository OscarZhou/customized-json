package logic

import (
	"errors"
	"reflect"
)

// Mapper is a
type Mapper interface {
	ToKeys() ([]string, error)
}

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
