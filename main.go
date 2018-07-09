package main

import (
	"customized-json/models"
	"fmt"
	"html/template"
	"net/http"
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

func indexHandler(w http.ResponseWriter, r *http.Request, jt models.JSONTemplate) {
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "fail to render index.html", http.StatusInternalServerError)
	}

	fmt.Println("----", t)
	t.Execute(w, jt)
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, jt models.JSONTemplate)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// c := models.NewConfig(templates)
		jt := models.JSONTemplate{
			Title: "First jt",
		}
		fn(w, r, jt)
	}
}

func main() {
	// c := models.NewConfig(templates)
	// err := c.OutputConfigFile("route")
	// if err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("/index", makeHandler(indexHandler))

	http.ListenAndServe(":7000", nil)
}
