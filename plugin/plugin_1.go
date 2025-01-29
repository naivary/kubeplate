package main

import (
	"strings"
	"text/template"
)

var FuncsTmpl = template.FuncMap{
	"toLower": strings.ToLower,
}
