package main

import (
	"customized-json/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
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

func indexHandler(ctx *gin.Context, jt models.JSONTemplate) {
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "fail to render index.html", nil)
		return
	}
	t.Execute(ctx.Writer, jt)
}

func makeHandler(fn func(ctx *gin.Context, jt models.JSONTemplate)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// c := models.NewConfig(templates)
		jt := models.JSONTemplate{
			Title: "First jt",
		}
		fn(ctx, jt)
	}
}

func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

func main() {
	// c := models.NewConfig(templates)
	// err := c.OutputConfigFile("route")
	// if err != nil {
	// 	panic(err)
	// }

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	// r.StaticFile("bootstrap.css", "./public/css/bootstrap.css")
	// r.StaticFile("main.css", "./public/css/main.css")
	// r.StaticFile("bootstrap.js", "./public/js/bootstrap.js")
	// r.StaticFile("jquery-3.1.1.js", "./public/js/jquery-3.1.1.js")

	r.StaticFS("/public", http.Dir("public"))

	r.Handle("GET", "/index", makeHandler(indexHandler))
	r.Run(":7000")
}
