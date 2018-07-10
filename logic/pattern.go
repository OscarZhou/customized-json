package logic

import (
	"errors"
	"reflect"
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

func (p *Pattern) Assemble(content string) error {

	// keyPairs := strings.Split(content, ",")
	// for _, v := range keyPairs {
	// 	pairs := strings.TrimSpace(v)
	// 	s := strings.Split(pairs, " ")
	// 	if len(s) != 2 {
	// 		return  errors.New("format is invalid")
	// 	}

	// 	modelStruct.KeyTypes = append(modelStruct.KeyTypes, KeyType{Key: s[0], Type: s[1]})
	// }
	// return modelStruct, nil

	// keys, err := p.ToKeys()
	// if err != nil {
	// 	return err
	// }

	// for _, v := range keys {

	// }
	return nil
}
