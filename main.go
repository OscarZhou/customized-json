package main

import (
	"customized-json/models"
)

var templates []models.Template

func init() {
	apiResources := []string{
		"Analysis",
		"AnalysisOutputs",
		"ExtraCurriculums",
		"Favourites",
		"FavouriteExtraCurriculums",
		"Imds",
		"Meshblocks",
		"MeshblockMappings",
		"Rankings",
		"RentPriceByTaAndCities",
		"SalePriceByTaAndCities",
		"Scholarships",
		"SchoolDeciles",
		"SchoolRolls",
		"SchoolRollCategories",
		"Schools",
		"SchoolDataByYears",
		"SchoolZones",
		"NZAcademics",
		"NZSchoolAcademics",
		"NZYearLevels",
		"NZCategories",
		"NZPerformances",
		"NZPerformanceValues",
		"DynamicFormQuestions",
		"SchoolDatazones",
		"SchoolImds",
	}
	t := models.Template{
		ServerTemplate: models.ServerTemplate{
			ProjectVersion: "v1",
			APIVersion:     "v1",
			ProxySchema:    "http",
			ProxyPass:      "127.0.0.1:8021",
			CustomConfigs: models.CustomConfig{
				"Cached":                  true,
				"CachedDurationsInSecond": 10,
				"Authentication":          false,
			},
		},
		APIs:      make(map[string][]string),
		Resources: apiResources,
	}
	templates = append(templates, t)

	t = models.Template{
		ServerTemplate: models.ServerTemplate{
			ProjectVersion: "v0.1",
			APIVersion:     "v1",
			ProxySchema:    "http",
			ProxyPass:      "127.0.0.1:8021",
			CustomConfigs: models.CustomConfig{
				"Cached":                  true,
				"CachedDurationsInSecond": 10,
				"Authentication":          true,
			},
		},
		APIs:      make(map[string][]string),
		Resources: apiResources,
	}
	templates = append(templates, t)

	t = models.Template{
		ServerTemplate: models.ServerTemplate{
			ProjectVersion: "",
			APIVersion:     "v1",
			ProxySchema:    "http",
			ProxyPass:      "127.0.0.1:8060",
			CustomConfigs: models.CustomConfig{
				"Cached":                  true,
				"CachedDurationsInSecond": 10,
				"Authentication":          false,
			},
		},
		APIs: make(map[string][]string),
		Resources: []string{
			"Activations",
			"Sms",
			"authorize",
			"token",
			"jwt",
			"jwtRegisters",
		},
	}
	templates = append(templates, t)

}

func main() {
	// http.HandleFunc("/index")

	c := models.NewConfig(templates)
	err := c.OutputConfigFile("route")
	if err != nil {
		panic(err)
	}

}
