package main

import (
	"customized-json/models"
	"customized-json/utils"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Define template
	var templates []models.ServerTemplate

	serverTemplate := models.ServerTemplate{
		ProjectVersion: "v0.4",
		APIVersion:     "v1",
		ProxySchema:    "http",
		ProxyPass:      "127.0.0.1:9100",
	}

	templates = append(templates, serverTemplate)

	// for _, v := range templates {

	// }

	// Load controller path to parse Controller
	controllerPath := `C:\Users\project\cms\controllers`

	fileList := []string{}

	err := filepath.Walk(controllerPath, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	f := `C:\Users\project\cms\controllers\article.go`
	methods, err := utils.ParseAPIMethod(f)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range methods {
		fmt.Println(v)
	}

	// LoadControllers

	//Generate route.json
	var servers []models.Server
	for _, template := range templates {
		for _, method := range methods {
			path := "/" + template.APIVersion + "/Articles"
			if method == "GETBYID" {
				path += "/:id"
			}
			s := models.Server{
				Method:        method,
				Path:          path,
				ProxyScheme:   template.ProxySchema,
				ProxyPass:     template.ProxyPass,
				ProxyPassPath: "/Articles",
				APIVersion:    template.APIVersion,
			}

			servers = append(servers, s)
		}
	}

	// cp, err := New()
	// if err != nil {
	// 	panic(err)
	// }

	// cp.Add("Method", "GET")
	// cp.Add("Path", "/v0.3/Accounts")
	// cp.Add("ProxyScheme", "http")
	// cp.Add("ProxyPass", "127.0.0.1:9100")
	// cp.Add("ProxyPassPath", "/Accounts")
	// cp.Add("APIVersion", "v0.3")

	// cp.Add("CustomConfigs", "/v0.3/Accounts")
	// cp.Add("Path", "/v0.3/Accounts")
	// fmt.Println(cp)

	// err = cp.CreateFile("test")
	// if err != nil {
	// 	panic(err)
	// }

	err = utils.CreateFile("test", servers)
	if err != nil {
		panic(err)
	}
}
