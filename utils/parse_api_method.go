package utils

import (
	"bufio"
	"os"
	"regexp"
)

func ParseAPIMethod(fileName string) ([]string, error) {
	var APIMethods []string
	registerAPIMethods := make(map[string]string)

	registerAPIMethods["GetByID"] = "GETBYID"
	registerAPIMethods["Get"] = "GET"
	registerAPIMethods["Post"] = "POST"
	registerAPIMethods["Patch"] = "PATCH"

	inFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		reg1 := regexp.MustCompile("func").MatchString
		if reg1(line) {
			method, err := ParseFuncName(line, registerAPIMethods)
			if err != nil {
				return nil, err
			}

			if method != "" {
				APIMethods = append(APIMethods, method)
			}
		}

	}

	return APIMethods, nil
}

func ParseFuncName(line string, funcName map[string]string) (string, error) {
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
