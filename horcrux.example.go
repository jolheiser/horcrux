// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"go.jolheiser.com/beaver"
)

const (
	exampleFile = "horcrux.example.yml"
	writeFile   = "config/config_default.go"
	tmpl        = `package config

func init() {
	defaultConfig = ` + "`" + `%s` + "`" + `
}
`
)

func main() {
	bytes, err := ioutil.ReadFile(exampleFile)
	if err != nil {
		beaver.Fatalf("Could not read from %s. Are you in the root directory of the project?", exampleFile)
	}

	data := fmt.Sprintf(tmpl, string(bytes))
	if err := ioutil.WriteFile(writeFile, []byte(data), os.ModePerm); err != nil {
		beaver.Fatalf("Could not write to %s.", writeFile)
	}
}
