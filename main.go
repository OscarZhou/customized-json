package main

import (
	"customized-json/models"
)

func main() {
	var templates []models.Template

	t := models.Template{
		ControllerPath: `C:\Users\project\ems\controllers`,
		ServerTemplate: models.ServerTemplate{
			ProjectVersion: "v0.4",
			APIVersion:     "v1",
			ProxySchema:    "http",
			ProxyPass:      "127.0.0.1:9100",
			CustomConfigs: models.CustomConfig{
				"Cached":                  true,
				"CachedDurationsInSecond": 10,
				"Authentication":          true,
			},
		},
		APIs: make(map[string][]string),
	}
	templates = append(templates, t)

	t = models.Template{}
	t = models.Template{
		ControllerPath: `C:\Users\project\cms\controllers`,
		ServerTemplate: models.ServerTemplate{
			ProjectVersion: "v0.4",
			APIVersion:     "v1",
			ProxySchema:    "http",
			ProxyPass:      "127.0.0.1:9101",
			CustomConfigs: models.CustomConfig{
				"Cached":                  true,
				"CachedDurationsInSecond": 10,
				"Authentication":          false,
			},
		},
		APIs: make(map[string][]string),
	}
	templates = append(templates, t)

	c := models.NewConfig(templates)
	err := c.OutputConfigFile("route")
	if err != nil {
		panic(err)
	}

}
