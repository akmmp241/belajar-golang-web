package belajar_golang_web

import (
	"html/template"
)

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))
