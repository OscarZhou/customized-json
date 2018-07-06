package main

import "html/template"

var pageTemplates = template.Must(template.ParseFiles("index.html"))

// func renderTemplate(w http.ResponseWriter, tmpl string, t *models.Config){
// 	err :=
// }
