package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Resources     []string
	Version       string
	ProxySchema   string
	ProxyPass     string
	ProxyVersion  string
	CustomConfigs CustomConfig
}

func SetRoute(ctx *gin.Context, r Route) {
	resources := ctx.PostForm("resource")
	fmt.Println(resources)
	resources = strings.Replace(resources, "\r\n", "", -1)
	resources = strings.Replace(resources, "\"", "", -1)
	resources = strings.Replace(resources, "\n", "", -1)
	resources = strings.Replace(resources, "\r", "", -1)
	resources = strings.Replace(resources, " ", "", -1)
	r.Resources = strings.Split(resources[:len(resources)-1], ",")

	version := ctx.PostForm("version")
	r.Version = version

	proxySchema := ctx.PostForm("proxy_schema")
	r.ProxySchema = proxySchema

	proxyPass := ctx.PostForm("proxy_pass")
	r.ProxyPass = proxyPass

	proxyVersion := ctx.PostForm("proxy_version")
	r.ProxyVersion = proxyVersion

	customConfig := ctx.PostForm("custom_config")
	json.Unmarshal([]byte(customConfig), &r.CustomConfigs)

	fmt.Println(r)

	t := Template{
		ServerTemplate: ServerTemplate{
			ProjectVersion: r.Version,
			APIVersion:     r.ProxyVersion,
			ProxySchema:    r.ProxySchema,
			ProxyPass:      r.ProxyPass,
			CustomConfigs:  r.CustomConfigs,
		},
		APIs:      make(map[string][]string),
		Resources: r.Resources,
	}

	templates := []Template{
		t,
	}

	c := NewConfig(templates)
	err := c.OutputConfigFile("newRoute")
	if err != nil {
		panic(err)
	}
	// keys, err := p.ToKeys()
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// // ctx.JSON(http.StatusOK, keys)
	// ctx.HTML(http.StatusOK, "add_pattern.html", gin.H{
	// 	"Pattern": keys,
	// })
	return
}

func GetRoute(ctx *gin.Context) {
	r := Route{}
	ctx.HTML(http.StatusOK, "view_route.html", gin.H{
		"Route": r,
	})
}
