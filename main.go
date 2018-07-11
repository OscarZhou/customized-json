package main

import (
	"customized-json/api"
	"customized-json/logic"
	"customized-json/models"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

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

func groupContainer(fn func(ctx *gin.Context, r models.Route)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := models.Route{}
		fn(ctx, r)
	}
}
func createModelHandler(ctx *gin.Context) {
	t, err := template.ParseFiles("templates/create_model.html")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "fail to render CreateModel.html", nil)
		return
	}
	t.Execute(ctx.Writer, nil)
}

func generateModelHandler(ctx *gin.Context) {
	var jsonTemplate models.JSONTemplate
	jsonTemplate.Title = ctx.Param("filename")
	jsonTemplate.Content = ctx.PostForm("content")

	if err := jsonTemplate.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, "error: %v", err)
		return
	}

	_, err := jsonTemplate.Parse()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "error: %v", err)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/ViewModel")
}

func viewModelHandler(ctx *gin.Context) {
	if err := filepath.Walk("files/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			jsonTemplate, err := models.LoadJSONTemplate(path)
			if err != nil {
				if jsonTemplate == nil {
					return nil
				}
				return err
			}

			modelStruct, err := jsonTemplate.Parse()
			if err != nil {
				return err
			}
			t, err := template.ParseFiles("templates/view_model.html")
			if err != nil {
				return err
			}
			t.Execute(ctx.Writer, modelStruct)
		}
		return nil
	}); err != nil {
		ctx.String(http.StatusInternalServerError, "error: %v", err)
		return
	}
}

func indexHandler(ctx *gin.Context, jt models.JSONTemplate) {
	t, err := template.ParseFiles("templates/create_model.html")
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

func mapContainer(fn func(ctx *gin.Context, p logic.Mapper)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, _ := logic.Default()
		m, _ := t.MapOut()
		fn(ctx, m)
	}
}

func assembleContainer(fn func(ctx *gin.Context, p logic.Assembler)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := make(logic.Pattern)
		fn(ctx, &p)
	}
}

func templateContainer(fn func(ctx *gin.Context, t logic.Templator)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, _ := logic.Default()
		fn(ctx, t)
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

	r.StaticFS("/public", http.Dir("public"))
	r.LoadHTMLGlob("templates/*")

	// r.GET("/Error", func(ctx *gin.Context ))}{
	// 	ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
	// 		"message":""
	// 	})
	// }

	r.GET("/CreateModel", createModelHandler)
	r.GET("/ViewModel", viewModelHandler)
	r.GET("/GenerateModel/:filename", generateModelHandler)

	r.GET("/AddPattern", mapContainer(api.GetPattern))
	r.GET("/GetTemplate", templateContainer(api.GetTemplate))
	r.POST("/SetPattern", assembleContainer(api.SetPattern))

	r.GET("/ViewRoute", models.GetRoute)
	r.POST("/SetRoute", groupContainer(models.SetRoute))

	// r.Handle("GET", "/CreateModel", makeHandler(indexHandler))
	r.Run(":7000")
}
