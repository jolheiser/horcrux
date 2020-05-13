package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// All template names will be expected to input/output based on the following:
// input: router/templates/<name>.tmpl
// output: router/<name>_tmpl.go
var templates = []string{
	"index",
	"info",
}

func main() {
	template.Must(template.New("name").Parse(""))
	for _, tmpl := range templates {
		data, err := ioutil.ReadFile(fmt.Sprintf("router/templates/%s.tmpl", tmpl))
		if err != nil {
			panic(err)
		}

		contents := fmt.Sprintf(routerTmpl, tmpl, tmpl, string(data))
		if err := ioutil.WriteFile(fmt.Sprintf("router/%s_tmpl.go", tmpl), []byte(contents), os.ModePerm); err != nil {
			panic(err)
		}
	}
}

var routerTmpl = `package router

import "html/template"

func init() {
	%sTmpl = template.Must(template.New("%s").Funcs(funcMap()).Parse(` + "`" + `%s` + "`" + `))
}
`
